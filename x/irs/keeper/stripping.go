package keeper

import (
	"encoding/json"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
	yieldaggregatorkeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"

	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
)

func (k Keeper) MintPtYtPair(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, depositAmount sdk.Coin) (math.Int, error) {
	// Send coins from sender to IRS vault account
	moduleAddr := types.GetVaultModuleAddress(pool)
	err := k.bankKeeper.SendCoins(ctx, sender, moduleAddr, sdk.Coins{depositAmount})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	ptDenom := types.PtDenom(pool)
	ytDenom := types.YtDenom(pool)

	ptAmount, err := k.CalculateMintPtAmount(ctx, pool, depositAmount)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	ptCoins := sdk.Coins{sdk.NewCoin(ptDenom, ptAmount)}

	contractAddr := sdk.MustAccAddressFromBech32(pool.StrategyContract)

	// Stake to strategy
	if depositAmount.Denom == pool.DepositDenom {
		wasmMsg := `{"stake":{}}`
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, moduleAddr, []byte(wasmMsg), sdk.Coins{depositAmount})
		if err != nil {
			return sdk.ZeroInt(), err
		}
	} else {
		return sdk.ZeroInt(), types.ErrInvalidDepositDenom
		// TODO: bug here, PT is calculated and minted regardless of denom
		msg, err := k.ExecuteVaultTransfer(ctx, moduleAddr, pool.StrategyContract, depositAmount)
		k.Logger(ctx).Info("transfer_memo " + msg.Memo)
		if err != nil {
			return sdk.ZeroInt(), err
		}
	}

	// mint PT
	// PT mint amount = usedUnderlying * (1-(strategyAmount)/interestSupply)
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, ptCoins)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, ptCoins)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	// mint YT
	// YT mint amount = usedUnderlying
	info := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	rate := sdk.MustNewDecFromStr(info.DepositDenomRate)
	if rate.IsZero() {
		return sdk.ZeroInt(), types.ErrZeroDepositRate
	}
	ytAmount := sdk.NewDecFromInt(depositAmount.Amount).Quo(rate).TruncateInt()

	ytCoins := sdk.Coins{sdk.NewCoin(ytDenom, ytAmount)}
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, ytCoins)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, ytCoins)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	return ptAmount, nil
}

// CalculateMintPtAmount calculates the amount of PT to be minted PT & YT pair
func (k Keeper) CalculateMintPtAmount(ctx sdk.Context, pool types.TranchePool, depositAmount sdk.Coin) (math.Int, error) {
	moduleAddr := types.GetVaultModuleAddress(pool)

	ytDenom := types.YtDenom(pool)
	ptDenom := types.PtDenom(pool)
	interestSupply := k.bankKeeper.GetSupply(ctx, ytDenom)
	ptSupply := k.bankKeeper.GetSupply(ctx, ptDenom)

	// Initial deposit
	if interestSupply.IsZero() {
		return depositAmount.Amount, nil
	}

	amountFromStrategy, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	info := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	rate := sdk.MustNewDecFromStr(info.DepositDenomRate)
	if rate.IsZero() {
		return sdk.ZeroInt(), types.ErrZeroDepositRate
	}
	utAmount := sdk.NewDecFromInt(depositAmount.Amount).Quo(rate).TruncateInt()

	// mint PT
	if ptSupply.IsPositive() && amountFromStrategy.GT(ptSupply.Amount) {
		// PT mint amount = usedUnderlying * (1-(strategyAmount-ptAmount)/interestSupply)
		ptAmount := utAmount.
			Sub(
				utAmount.
					Mul(amountFromStrategy.Sub(ptSupply.Amount)).
					Quo(interestSupply.Amount),
			)
		return ptAmount, nil
	}

	return utAmount, nil
}

// RedeemPtYtPair redeems Pt and Yt pair and can be executed before maturity
// The ratio between Pt Supply : Yt Supply and Pt / Yt amount redeemed should be same
func (k Keeper) RedeemPtYtPair(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, redeemAmount math.Int, maxPtYtIns sdk.Coins) error {
	moduleAddr := types.GetVaultModuleAddress(pool)

	if redeemAmount.IsZero() {
		return types.ErrZeroAmount
	}

	amountFromStrategy, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return err
	}
	if amountFromStrategy.IsZero() {
		return types.ErrZeroAmount
	}

	requiredPt, requiredYt, err := k.CalculateRedeemRequiredPtAndYtAmount(ctx, pool, redeemAmount)
	if err != nil {
		return err
	}

	coins := sdk.Coins{}
	requiredPtYt := coins.Add(requiredPt).Add(requiredYt)
	if !maxPtYtIns.IsAllGTE(requiredPtYt) {
		return types.ErrInSufficientTokenInMaxs
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, requiredPtYt)
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, requiredPtYt)
	if err != nil {
		return err
	}

	return k.UnstakeFromStrategy(ctx, moduleAddr, sender.String(), pool.StrategyContract, redeemAmount)
}

func (k Keeper) CalculateRedeemRequiredPtAndYtAmount(ctx sdk.Context, pool types.TranchePool, redeemAmount math.Int) (sdk.Coin, sdk.Coin, error) {
	moduleAddr := types.GetVaultModuleAddress(pool)

	amountFromStrategy, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}
	if amountFromStrategy.IsZero() {
		return sdk.Coin{}, sdk.Coin{}, types.ErrZeroAmount
	}
	ptDenom := types.PtDenom(pool)
	ytDenom := types.YtDenom(pool)
	ptSupply := k.bankKeeper.GetSupply(ctx, ptDenom)
	ytSupply := k.bankKeeper.GetSupply(ctx, ytDenom)
	info := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	rate := sdk.MustNewDecFromStr(info.DepositDenomRate)
	if rate.IsZero() {
		return sdk.Coin{}, sdk.Coin{}, types.ErrZeroDepositRate
	}
	utAmount := sdk.NewDecFromInt(redeemAmount).Quo(rate).TruncateInt()
	requiredPtAmount := ptSupply.Amount.Mul(utAmount).Quo(amountFromStrategy)
	requiredYtAmount := ytSupply.Amount.Mul(utAmount).Quo(amountFromStrategy)
	requiredPt := sdk.NewCoin(ptDenom, requiredPtAmount)
	requiredYt := sdk.NewCoin(ytDenom, requiredYtAmount)
	return requiredPt, requiredYt, nil
}

func (k Keeper) CalculateRedeemAmount(ctx sdk.Context, pool types.TranchePool, tokenIn sdk.Coin) (sdk.Coin, sdk.Coin, error) {
	ptDenom := types.PtDenom(pool)
	ytDenom := types.YtDenom(pool)
	redeemDenom := pool.DepositDenom

	moduleAddr := types.GetVaultModuleAddress(pool)
	amountFromStrategy, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return sdk.Coin{}, sdk.Coin{}, err
	}
	if amountFromStrategy.IsZero() {
		return sdk.Coin{}, sdk.Coin{}, types.ErrZeroAmount
	}

	ptSupply := k.bankKeeper.GetSupply(ctx, ptDenom)
	if ptSupply.IsZero() {
		return sdk.Coin{}, sdk.Coin{}, types.ErrSupplyNotFound
	}
	ytSupply := k.bankKeeper.GetSupply(ctx, ytDenom)
	if ytSupply.IsZero() {
		return sdk.Coin{}, sdk.Coin{}, types.ErrSupplyNotFound
	}
	var utAmount math.Int
	var requiredCoin sdk.Coin

	if tokenIn.Denom == ptDenom {
		ptSupply := k.bankKeeper.GetSupply(ctx, ptDenom)
		utAmount = amountFromStrategy.Mul(tokenIn.Amount).Quo(ptSupply.Amount)
		requiredAmount := tokenIn.Amount.Mul(ytSupply.Amount).Quo(ptSupply.Amount)
		requiredCoin = sdk.NewCoin(ytDenom, requiredAmount)
	} else if tokenIn.Denom == ytDenom {
		utAmount = amountFromStrategy.Mul(tokenIn.Amount).Quo(ytSupply.Amount)
		requiredAmount := tokenIn.Amount.Mul(ptSupply.Amount).Quo(ytSupply.Amount)
		requiredCoin = sdk.NewCoin(ptDenom, requiredAmount)
	}

	info := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	rate := sdk.MustNewDecFromStr(info.DepositDenomRate)
	redeemAmount := sdk.NewDecFromInt(utAmount).Mul(rate).TruncateInt()

	redeemCoin := sdk.NewCoin(redeemDenom, redeemAmount)
	return redeemCoin, requiredCoin, nil
}

func (k Keeper) RedeemPtAtMaturity(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, ptAmount sdk.Coin) error {
	if uint64(ctx.BlockTime().Unix()) < pool.StartTime+pool.Maturity {
		return types.ErrTrancheNotMatured
	}
	ptDenom := types.PtDenom(pool)
	if ptDenom != ptAmount.Denom {
		return types.ErrInvalidPtDenom
	}
	if ptAmount.IsZero() {
		return types.ErrZeroAmount
	}

	moduleAddr := types.GetVaultModuleAddress(pool)
	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{ptAmount})
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{ptAmount})
	if err != nil {
		return err
	}

	info := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	rate := sdk.MustNewDecFromStr(info.DepositDenomRate)
	redeemAmount := sdk.NewDecFromInt(ptAmount.Amount).Mul(rate).TruncateInt()

	return k.UnstakeFromStrategy(ctx, moduleAddr, sender.String(), pool.StrategyContract, redeemAmount)
}

func (k Keeper) RedeemYtAtMaturity(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, ytAmount sdk.Coin) error {
	if uint64(ctx.BlockTime().Unix()) < pool.StartTime+pool.Maturity {
		return types.ErrTrancheNotMatured
	}
	moduleAddr := types.GetVaultModuleAddress(pool)
	redeemAmount, err := k.CalculateRedeemYtAmount(ctx, pool, ytAmount)
	if err != nil {
		return err
	}
	if redeemAmount.IsZero() {
		return nil
	}

	err = k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{ytAmount})
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{ytAmount})
	if err != nil {
		return err
	}

	return k.UnstakeFromStrategy(ctx, moduleAddr, sender.String(), pool.StrategyContract, redeemAmount)
}

func (k Keeper) CalculateRedeemYtAmount(ctx sdk.Context, pool types.TranchePool, ytAmount sdk.Coin) (math.Int, error) {
	ptDenom := types.PtDenom(pool)
	ytDenom := types.YtDenom(pool)
	if ytDenom != ytAmount.Denom {
		return sdk.ZeroInt(), types.ErrInvalidYtDenom
	}
	if ytAmount.IsZero() {
		return sdk.ZeroInt(), types.ErrZeroAmount
	}
	ptSupply := k.bankKeeper.GetSupply(ctx, ptDenom)
	ytSupply := k.bankKeeper.GetSupply(ctx, ytDenom)
	if ytSupply.IsZero() {
		return sdk.ZeroInt(), types.ErrSupplyNotFound
	}

	moduleAddr := types.GetVaultModuleAddress(pool)
	vaultAmount, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	utAmount := vaultAmount.Sub(ptSupply.Amount).Mul(ytAmount.Amount).Quo(ytSupply.Amount)
	info := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	rate := sdk.MustNewDecFromStr(info.DepositDenomRate)
	redeemAmount := sdk.NewDecFromInt(utAmount).Mul(rate).TruncateInt()
	return redeemAmount, nil
}

func (k Keeper) ExecuteVaultTransfer(ctx sdk.Context, moduleAddr sdk.AccAddress, strategyContract string, stakeCoin sdk.Coin) (*ibctypes.MsgTransfer, error) {
	contractAddr := sdk.MustAccAddressFromBech32(strategyContract)
	info := k.GetStrategyDepositInfo(ctx, strategyContract)
	params, err := k.YieldaggregatorKeeper.GetParams(ctx)
	if err != nil {
		return nil, err
	}
	ibcTransferTimeoutNanos := params.IbcTransferTimeoutNanos
	timeoutTimestamp := uint64(ctx.BlockTime().UnixNano()) + ibcTransferTimeoutNanos
	denomInfo := k.YieldaggregatorKeeper.GetDenomInfo(ctx, stakeCoin.Denom)
	symbolInfo := k.YieldaggregatorKeeper.GetSymbolInfo(ctx, denomInfo.Symbol)
	tarChannels := []yieldaggregatortypes.TransferChannel{}
	for _, channel := range symbolInfo.Channels {
		if channel.RecvChainId == info.TargetChainId {
			tarChannels = []yieldaggregatortypes.TransferChannel{channel}
			break
		}
	}
	// increase vault pending deposit
	k.recordsKeeper.IncreaseVaultPendingDeposit(ctx, 0, stakeCoin.Amount)

	// calculate transfer route and execute the transfer
	transferRoute := yieldaggregatorkeeper.CalculateTransferRoute(denomInfo.Channels, tarChannels)
	initialReceiver, metadata := k.YieldaggregatorKeeper.ComposePacketForwardMetadata(ctx, transferRoute, info.TargetChainAddr)
	memo, err := json.Marshal(metadata)
	if err != nil {
		return nil, err
	}
	if metadata == nil {
		memo = []byte{}
	}
	msg := ibctypes.NewMsgTransfer(
		ibctransfertypes.PortID,
		transferRoute[0].ChannelId,
		stakeCoin,
		moduleAddr.String(),
		initialReceiver,
		clienttypes.Height{},
		timeoutTimestamp,
		string(memo),
	)
	err = k.recordsKeeper.VaultTransfer(ctx, 0, contractAddr, msg)
	return msg, err
}
