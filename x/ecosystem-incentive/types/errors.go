package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidTotalWeight          = sdkerrors.Register(ModuleName, 1, "total weight in an incentive-unit is invalid")
	ErrSubjectsWeightsNumUnmatched = sdkerrors.Register(ModuleName, 2, "the numbers of subjects and weights must be same")
	ErrRegisteredIncentiveId       = sdkerrors.Register(ModuleName, 3, "the incentive id is already used")
)
