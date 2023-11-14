package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/irs/types"
)

func (k msgServer) RegisterInterestRateSwapVault(goCtx context.Context, msg *types.MsgRegisterInterestRateSwapVault) (*types.MsgRegisterInterestRateSwapVaultResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx

	return &types.MsgRegisterInterestRateSwapVaultResponse{}, nil
}
