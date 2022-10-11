package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidTotalWeight           = sdkerrors.Register(ModuleName, 1, "total weight in an incentive_unit is invalid")
	ErrSubjectsWeightsNumUnmatched  = sdkerrors.Register(ModuleName, 2, "the numbers of subjects and weights must be same")
	ErrRegisteredIncentiveId        = sdkerrors.Register(ModuleName, 3, "the incentive id is already used")
	ErrRewardNotExists              = sdkerrors.Register(ModuleName, 4, "the reward for the subject doesn't exist")
	ErrDenomRewardNotExists         = sdkerrors.Register(ModuleName, 5, "the reward of the denom for the subject doesn't exist")
	ErrCoinAmount                   = sdkerrors.Register(ModuleName, 6, "must be 0 amount after sending token")
	ErrRecordedNftId                = sdkerrors.Register(ModuleName, 7, "the nft_id is already recorded")
	ErrNotRegisteredIncentiveUnitId = sdkerrors.Register(ModuleName, 8, "the incentive_unit_id is not registered")
	ErrNftIdForFrontendDoesntExist  = sdkerrors.Register(ModuleName, 9, "the nft_id is not recorded")
	ErrRewardRateNotFound           = sdkerrors.Register(ModuleName, 10, "the reward rate in the params was not found")
	ErrAddressNotHaveReward         = sdkerrors.Register(ModuleName, 11, "the address doesn't have any rewards")
	ErrUnknownMemoVersion           = sdkerrors.Register(ModuleName, 12, "the version in the memo inputs is unknown")
	ErrInvalidIncentiveUnitIdLen    = sdkerrors.Register(ModuleName, 13, "the length of incentive unit id is invalid")
)
