package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewReward(subject sdk.AccAddress, rewards sdk.Coins) Reward {
	return Reward{
		SubjectAddr: subject.Bytes(),
		Rewards:     rewards,
	}
}
