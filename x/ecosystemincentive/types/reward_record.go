package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewRewardRecord(address string, rewards sdk.Coins) RewardRecord {
	return RewardRecord{
		Address: address,
		Rewards: rewards,
	}
}
