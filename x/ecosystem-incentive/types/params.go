package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	nftmarkettypes "github.com/UnUniFi/chain/x/nftmarket/types"
)

var (
	DafaultRewardParams = []*RewardParams{
		{
			ModuleName: nftmarkettypes.ModuleName,
			RewardRate: []RewardRate{
				{
					RewardType: RewardType_NFTMARKET_FRONTEND,
					Rate:       sdk.MustNewDecFromStr("0.5"),
				},
			},
		},
		{
			ModuleName: nftmarkettypes.ModuleName,
			RewardRate: []RewardRate{
				{
					RewardType: RewardType_STAKERS,
					Rate:       sdk.MustNewDecFromStr("0.5"),
				},
			},
		},
	}
	DefaultMaxIncentiveUnitIdLen   uint64 = 128
	DefaultMaxSubjectInfoNumInUnit uint64 = 15
	KeyRewardParams                       = []byte("RewardParams")
	KeyMaxIncentiveUnitIdLen              = []byte("MaxIncentiveUnitId")
	KeyMaxSubjectInfoNumInUnit            = []byte("MaxSubjectInfoNumInUnit")
)

// ParamKeyTable the param key table for launch module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return DefaultParams()
}

// DefaultParams returns a default set of parameters
func DefaultParams() Params {
	return Params{
		RewardParams:            DafaultRewardParams,
		MaxIncentiveUnitIdLen:   DefaultMaxIncentiveUnitIdLen,
		MaxSubjectInfoNumInUnit: DefaultMaxSubjectInfoNumInUnit,
	}
}

// ParamSetPairs get the params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyRewardParams, &p.RewardParams, validateRewardParams),
		paramtypes.NewParamSetPair(KeyMaxIncentiveUnitIdLen, &p.MaxIncentiveUnitIdLen, validateMaxIncentiveUnitIdLen),
		paramtypes.NewParamSetPair(KeyMaxSubjectInfoNumInUnit, &p.MaxSubjectInfoNumInUnit, validateMaxSubjectInfoNumInUnit),
	}
}

// Validate validates the set of params
func (p Params) Validate() error {

	if err := validateRewardParams(p.RewardParams); err != nil {
		return err
	}

	if err := validateMaxIncentiveUnitIdLen(p.MaxIncentiveUnitIdLen); err != nil {
		return err
	}

	if err := validateMaxSubjectInfoNumInUnit(p.MaxSubjectInfoNumInUnit); err != nil {
		return err
	}

	return nil
}

func validateRewardParams(i interface{}) error {
	rewardParams, ok := i.([]*RewardParams)
	if !ok {
		return fmt.Errorf("invalid paramter type: %T", i)
	}

	for _, rewardParam := range rewardParams {
		for _, rate := range rewardParam.RewardRate {
			if rate.Rate.GT(sdk.OneDec()) {
				return fmt.Errorf("each reward rate must be less than 1 dec")
			}

			if rate.Rate.IsNegative() {
				return fmt.Errorf("each reward rate must be positive")
			}
		}
	}

	return nil
}

func validateMaxIncentiveUnitIdLen(i interface{}) error {
	maxIncentiveUnitIdLen, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	if maxIncentiveUnitIdLen < 0 {
		return fmt.Errorf("maxIncentiveUnitIdLen should be positive")
	}

	return nil
}

func validateMaxSubjectInfoNumInUnit(i interface{}) error {
	maxSubjectInfoNumInUnit, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter types: %T", i)
	}

	if maxSubjectInfoNumInUnit < 0 {
		return fmt.Errorf("maxSubjectInfoNumInUnit should be positive")
	}

	return nil
}
