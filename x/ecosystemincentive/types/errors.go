package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrRewardNotExists          = sdkerrors.Register(ModuleName, 1, "the reward for the subject doesn't exist")
	ErrDenomRewardNotExists     = sdkerrors.Register(ModuleName, 2, "the reward of the denom for the subject doesn't exist")
	ErrCoinAmount               = sdkerrors.Register(ModuleName, 3, "must be 0 amount after sending token")
	ErrRecordedNftId            = sdkerrors.Register(ModuleName, 4, "the nft_id is already recorded")
	ErrNotRegisteredRecipient   = sdkerrors.Register(ModuleName, 5, "the recipient is not registered")
	ErrRecipientByNftIdNotExist = sdkerrors.Register(ModuleName, 6, "the recipient by nft_id doesn't exist")
	ErrRewardRateNotFound       = sdkerrors.Register(ModuleName, 7, "the reward rate in the params was not found")
	ErrAddressNotHaveReward     = sdkerrors.Register(ModuleName, 8, "the address doesn't have any rewards")
	ErrRewardExceedsFee         = sdkerrors.Register(ModuleName, 9, "the total reward exceeds the fee")
	ErrRewardRateIsZero         = sdkerrors.Register(ModuleName, 16, "the reward rate is set zero for")
)
