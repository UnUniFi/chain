package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) RegisterInterestRateSwapVault(goCtx context.Context, msg *types.MsgRegisterInterestRateSwapVault) (*types.MsgRegisterInterestRateSwapVaultResponse, error) {
	if k.authority != msg.Sender {
		return nil, sdkerrors.ErrUnauthorized
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	params, err := k.GetParams(ctx)
	if err != nil {
		return nil, err
	}

	info := k.GetStrategyDepositInfo(ctx, msg.StrategyContract)

	// register IRS vault
	k.SetVault(ctx, types.InterestRateSwapVault{
		StrategyContract: msg.StrategyContract,
		Denom:            info.Denom,
		DepositDenom:     info.DepositDenom,
		Name:             msg.Name,
		Description:      msg.Description,
		MaxMaturity:      msg.MaxMaturity,
		Cycle:            msg.Cycle,
		LastTrancheTime:  uint64(ctx.BlockTime().Unix()),
	})

	// register first tranche pool for the vault
	k.SetTranchePool(ctx, types.TranchePool{
		Id:               k.GetLastTrancheId(ctx) + 1,
		StrategyContract: msg.StrategyContract,
		Denom:            info.Denom,
		DepositDenom:     info.DepositDenom,
		StartTime:        uint64(ctx.BlockTime().Unix()),
		Maturity:         msg.MaxMaturity,
		SwapFee:          params.TradeFeeRate,
		ExitFee:          sdk.ZeroDec(),
		TotalShares:      sdk.Coin{},
		PoolAssets:       sdk.Coins{},
	})

	return &types.MsgRegisterInterestRateSwapVaultResponse{}, nil
}
