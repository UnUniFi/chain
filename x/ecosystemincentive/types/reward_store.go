package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewRewardStore(address string, rewards sdk.Coins) RewardStore {
	return RewardStore{
		Address: address,
		Rewards: rewards,
	}
}
