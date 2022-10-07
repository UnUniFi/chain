package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	ununifitypes "github.com/UnUniFi/chain/types"
)

func NewRewardStore(subject ununifitypes.StringAccAddress, rewards sdk.Coins) RewardStore {
	return RewardStore{
		SubjectAddr: subject,
		Rewards:     rewards,
	}
}
