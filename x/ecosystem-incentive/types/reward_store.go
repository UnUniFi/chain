package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewRewardStore(subject string, rewards sdk.Coins) RewardStore {
	return RewardStore{
		SubjectAddr: subject,
		Rewards:     rewards,
	}
}
