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

func (k Keeper) MintPtYtPair(ctx sdk.Context, sender sdk.AccAddress, strategyContract string, maturity uint64, underlyingAmount sdk.Coin) error {
	// Send coins from sender to IRS vault account
	moduleAddr := types.GetVaultModuleAddress(strategyContract, maturity)
	err := k.bankKeeper.SendCoins(ctx, sender, moduleAddr, sdk.Coins{underlyingAmount})
	if err != nil {
		return err
	}

	ptDenom := types.PtDenom(strategyContract, maturity)
	ytDenom := types.YtDenom(strategyContract, maturity)
	contractAddr := sdk.MustAccAddressFromBech32(strategyContract)
	depositInfo := k.GetStrategyDepositInfo(ctx, strategyContract)
	interestSupply := k.bankKeeper.GetSupply(ctx, ytDenom)

	amountFromStrategy, err := k.GetAmountFromStrategy(ctx, moduleAddr, strategyContract)
	if err != nil {
		return err
	}

	// Stake to strategy
	if underlyingAmount.Denom == depositInfo.Denom {
		wasmMsg := `{"stake":{}}`
		_, err := k.wasmKeeper.Execute(ctx, contractAddr, moduleAddr, []byte(wasmMsg), sdk.Coins{underlyingAmount})
		if err != nil {
			return err
		}
	} else {
		msg, err := k.ExecuteVaultTransfer(ctx, moduleAddr, strategyContract, underlyingAmount)
		k.Logger(ctx).Info("transfer_memo " + msg.Memo)
		if err != nil {
			return err
		}
	}

	// mint PT
	// PT mint amount - usedUnderlying * (1-(strategyAmount)/interestSupply)
	ptAmount := underlyingAmount.Amount.
		Sub(underlyingAmount.Amount.Mul(amountFromStrategy).Quo(interestSupply.Amount))
	ptCoins := sdk.Coins{sdk.NewCoin(ptDenom, ptAmount)}
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, ptCoins)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, ptCoins)
	if err != nil {
		return err
	}

	// mint YT
	// YT mint amount - usedUnderlying
	ytCoins := sdk.Coins{sdk.NewCoin(ytDenom, underlyingAmount.Amount)}
	err = k.bankKeeper.MintCoins(ctx, types.ModuleName, ytCoins)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sender, ytCoins)
	if err != nil {
		return err
	}
	return nil
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
