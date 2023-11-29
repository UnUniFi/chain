package keeper

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
	yieldaggregatorkeeper "github.com/UnUniFi/chain/x/yieldaggregator/keeper"
	yieldaggregatortypes "github.com/UnUniFi/chain/x/yieldaggregator/types"

	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	ibctypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
)

func (k Keeper) MintPtYtPair(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, underlyingAmount sdk.Coin) (sdk.Int, error) {
	// Send coins from sender to IRS vault account
	moduleAddr := types.GetVaultModuleAddress(pool)
	err := k.bankKeeper.SendCoins(ctx, sender, moduleAddr, sdk.Coins{underlyingAmount})
	if err != nil {
		return sdk.ZeroInt(), err
	}

	ptDenom := types.PtDenom(pool)
	ytDenom := types.YtDenom(pool)
	contractAddr := sdk.MustAccAddressFromBech32(pool.StrategyContract)
	depositInfo := k.GetStrategyDepositInfo(ctx, pool.StrategyContract)
	interestSupply := k.bankKeeper.GetSupply(ctx, ytDenom)

	amountFromStrategy, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	// Stake to strategy
	if underlyingAmount.Denom == depositInfo.Denom {
		wasmMsg := `{"stake":{}}`
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, moduleAddr, []byte(wasmMsg), sdk.Coins{underlyingAmount})
		if err != nil {
			return sdk.ZeroInt(), err
		}
	} else {
		msg, err := k.ExecuteVaultTransfer(ctx, moduleAddr, pool.StrategyContract, underlyingAmount)
		k.Logger(ctx).Info("transfer_memo " + msg.Memo)
		if err != nil {
			return sdk.ZeroInt(), err
		}
	}

	// mint PT
	// PT mint amount - usedUnderlying * (1-(strategyAmount)/interestSupply)
	ptAmount := underlyingAmount.Amount.
		Sub(underlyingAmount.Amount.Mul(amountFromStrategy).Quo(interestSupply.Amount))
	ptCoins := sdk.Coins{sdk.NewCoin(ptDenom, ptAmount)}
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, ptCoins)
	if err != nil {
		return sdk.ZeroInt(), err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, ptCoins)
	if err != nil {
		return sdk.ZeroInt(), err
	}

	// mint YT
	// YT mint amount - usedUnderlying
	ytCoins := sdk.Coins{sdk.NewCoin(ytDenom, underlyingAmount.Amount)}
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

// RedeemPtYtPair redeems Pt and Yt pair and can be executed before maturity
// The ratio between Pt Supply : Yt Supply and Pt / Yt amount redeemed should be same
func (k Keeper) RedeemPtYtPair(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, redeemUt sdk.Int, maxPtYtIns sdk.Coins) error {
	moduleAddr := types.GetVaultModuleAddress(pool)

	amountFromStrategy, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return err
	}

	ptDenom := types.PtDenom(pool)
	ytDenom := types.YtDenom(pool)
	ptSupply := k.bankKeeper.GetSupply(ctx, ptDenom)
	ytSupply := k.bankKeeper.GetSupply(ctx, ytDenom)

	requiredPtAmount := ptSupply.Amount.Mul(redeemUt).Quo(amountFromStrategy)
	requiredYtAmount := ytSupply.Amount.Mul(redeemUt).Quo(amountFromStrategy)

	coins := sdk.Coins{}
	requiredPtYt := coins.Add(sdk.NewCoin(ptDenom, requiredPtAmount)).Add(sdk.NewCoin(ytDenom, requiredYtAmount))
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

	return k.UnstakeFromStrategy(ctx, moduleAddr, sender.String(), pool.StrategyContract, redeemUt)
}

func (k Keeper) RedeemPtAtMaturity(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, ptAmount sdk.Coin) error {
	moduleAddr := types.GetVaultModuleAddress(pool)
	if uint64(ctx.BlockTime().Unix()) < pool.StartTime+pool.Maturity {
		return types.ErrVaultNotMatured
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{ptAmount})
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{ptAmount})
	if err != nil {
		return err
	}

	ptDenom := types.PtDenom(pool)
	if ptDenom != ptAmount.Denom {
		return types.ErrInvalidPtDenom
	}
	if ptAmount.IsZero() {
		return nil
	}
	return k.UnstakeFromStrategy(ctx, moduleAddr, sender.String(), pool.StrategyContract, ptAmount.Amount)
}

func (k Keeper) RedeemYtAtMaturity(ctx sdk.Context, sender sdk.AccAddress, pool types.TranchePool, ytAmount sdk.Coin) error {
	moduleAddr := types.GetVaultModuleAddress(pool)
	if uint64(ctx.BlockTime().Unix()) < pool.StartTime+pool.Maturity {
		return types.ErrVaultNotMatured
	}

	err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.Coins{ytAmount})
	if err != nil {
		return err
	}

	err = k.bankKeeper.BurnCoins(ctx, types.ModuleName, sdk.Coins{ytAmount})
	if err != nil {
		return err
	}

	ptDenom := types.PtDenom(pool)
	ytDenom := types.YtDenom(pool)
	if ytDenom != ytAmount.Denom {
		return types.ErrInvalidYtDenom
	}

	ptSupply := k.bankKeeper.GetSupply(ctx, ptDenom)
	ytSupply := k.bankKeeper.GetSupply(ctx, ytDenom)
	vaultAmount, err := k.GetAmountFromStrategy(ctx, moduleAddr, pool.StrategyContract)
	if err != nil {
		return err
	}

	if ytAmount.IsZero() || ytSupply.IsZero() {
		return nil
	}

	redeemAmount := vaultAmount.Sub(ptSupply.Amount).Mul(ytAmount.Amount).Quo(ytSupply.Amount)
	if redeemAmount.IsZero() {
		return nil
	}

	return k.UnstakeFromStrategy(ctx, moduleAddr, sender.String(), pool.StrategyContract, redeemAmount)
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
