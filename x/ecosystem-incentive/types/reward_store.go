package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewReward(subject sdk.AccAddress, rewards sdk.Coins) RewardStore {
	return RewardStore{
		SubjectAddr: subject.Bytes(),
		Rewards:     rewards,
	}
}
