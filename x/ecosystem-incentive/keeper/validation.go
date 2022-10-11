package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/UnUniFi/chain/x/ecosystem-incentive/types"
)

// IncentiveUnitIdLenValidation checks MaxIncentiveUnitId validity
func (k Keeper) IncentiveUnitIdLenValidation(ctx sdk.Context, incentiveUnitId string) error {
	params := k.GetParams(ctx)

	if len(incentiveUnitId) > int(params.MaxIncentiveUnitIdLen) || len(incentiveUnitId) == 0 {
		return types.ErrInvalidIncentiveUnitIdLen
	}

	return nil
}
