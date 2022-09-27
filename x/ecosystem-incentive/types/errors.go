package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidTotalWeight          = sdkerrors.Register(ModuleName, 1, "total weight in an incentive-unit is invalid")
	ErrSubjectsWeightsNumUnmatched = sdkerrors.Register(ModuleName, 2, "the numbers of subjects and weights must be same")
	ErrRegisteredIncentiveId       = sdkerrors.Register(ModuleName, 3, "the incentive id is already used")
	ErrRewardNotExists             = sdkerrors.Register(ModuleName, 4, "the reward for the subject doesn't exist")
	ErrDenomRewardNotExists        = sdkerrors.Register(ModuleName, 5, "the reward of the denom for the subject doesn't exist")
	ErrCoinAmount                  = sdkerrors.Register(ModuleName, 6, "must be 0 amount after sending token")
)
