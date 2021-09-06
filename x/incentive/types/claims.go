package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	CdpMintingClaimType = "cdp_minting"
	BondDenom           = "ujsmn"
)

// Claim is an interface for handling common claim actions
type Claim interface {
	GetOwner() sdk.AccAddress
	GetReward() sdk.Coin
	GetType() string
}

// Claims is a slice of Claim
type Claims []Claim

// GetType returns the claim type, used to identify auctions in event attributes
func (c BaseClaim) GetType() string { return "base" }

// Validate performs a basic check of a BaseClaim fields
func (c BaseClaim) Validate() error {
	if c.Owner.AccAddress().Empty() {
		return errors.New("claim owner cannot be empty")
	}
	if !c.Reward.IsValid() {
		return fmt.Errorf("invalid reward amount: %s", c.Reward)
	}
	return nil
}

// GetType returns the claim type, used to identify auctions in event attributes
func (c BaseMultiClaim) GetType() string { return "base" }

// Validate performs a basic check of a BaseClaim fields
func (c BaseMultiClaim) Validate() error {
	if c.Owner.AccAddress().Empty() {
		return errors.New("claim owner cannot be empty")
	}
	if !sdk.Coins(c.Reward).IsValid() {
		return fmt.Errorf("invalid reward amount: %s", c.Reward)
	}
	return nil
}

// -------------- Custom Claim Types --------------

// NewCdpMintingClaim returns a new CdpMintingClaim
func NewCdpMintingClaim(owner sdk.AccAddress, reward sdk.Coin, rewardIndexes RewardIndexes) CdpMintingClaim {
	return CdpMintingClaim{
		BaseClaim: &BaseClaim{
			Owner:  owner.Bytes(),
			Reward: reward,
		},
		RewardIndexes: rewardIndexes,
	}
}

// GetType returns the claim's type
func (c CdpMintingClaim) GetType() string { return CdpMintingClaimType }

// GetReward returns the claim's reward coin
func (c CdpMintingClaim) GetReward() sdk.Coin { return c.Reward }

// GetOwner returns the claim's owner
func (c CdpMintingClaim) GetOwner() sdk.AccAddress {
	return c.Owner.AccAddress()
}

// Validate performs a basic check of a Claim fields
func (c CdpMintingClaim) Validate() error {
	if err := RewardIndexes(c.RewardIndexes).Validate(); err != nil {
		return err
	}

	return c.BaseClaim.Validate()
}

// HasRewardIndex check if a claim has a reward index for the input collateral type
func (c CdpMintingClaim) HasRewardIndex(collateralType string) (int64, bool) {
	for index, ri := range c.RewardIndexes {
		if ri.CollateralType == collateralType {
			return int64(index), true
		}
	}
	return 0, false
}

// CdpMintingClaims slice of CdpMintingClaim
type CdpMintingClaims []CdpMintingClaim

// Validate checks if all the claims are valid and there are no duplicated
// entries.
func (cs CdpMintingClaims) Validate() error {
	for _, c := range cs {
		if err := c.Validate(); err != nil {
			return err
		}
	}

	return nil
}

// ---------------------- Reward periods are used by the params ----------------------

// MultiRewardPeriod supports multiple reward types
type MultiRewardPeriod struct {
	Active           bool      `json:"active" yaml:"active"`
	CollateralType   string    `json:"collateral_type" yaml:"collateral_type"`
	Start            time.Time `json:"start" yaml:"start"`
	End              time.Time `json:"end" yaml:"end"`
	RewardsPerSecond sdk.Coins `json:"rewards_per_second" yaml:"rewards_per_second"` // per second reward payouts
}

// String implements fmt.Stringer
func (mrp MultiRewardPeriod) String() string {
	return fmt.Sprintf(`Reward Period:
	Collateral Type: %s,
	Start: %s,
	End: %s,
	Rewards Per Second: %s,
	Active %t,
	`, mrp.CollateralType, mrp.Start, mrp.End, mrp.RewardsPerSecond, mrp.Active)
}

// NewMultiRewardPeriod returns a new MultiRewardPeriod
func NewMultiRewardPeriod(active bool, collateralType string, start time.Time, end time.Time, reward sdk.Coins) MultiRewardPeriod {
	return MultiRewardPeriod{
		Active:           active,
		CollateralType:   collateralType,
		Start:            start,
		End:              end,
		RewardsPerSecond: reward,
	}
}

// Validate performs a basic check of a MultiRewardPeriod.
func (mrp MultiRewardPeriod) Validate() error {
	if mrp.Start.IsZero() {
		return errors.New("reward period start time cannot be 0")
	}
	if mrp.End.IsZero() {
		return errors.New("reward period end time cannot be 0")
	}
	if mrp.Start.After(mrp.End) {
		return fmt.Errorf("end period time %s cannot be before start time %s", mrp.End, mrp.Start)
	}
	if !mrp.RewardsPerSecond.IsValid() {
		return fmt.Errorf("invalid reward amount: %s", mrp.RewardsPerSecond)
	}
	if strings.TrimSpace(mrp.CollateralType) == "" {
		return fmt.Errorf("reward period collateral type cannot be blank: %s", mrp)
	}
	return nil
}

// MultiRewardPeriods array of MultiRewardPeriod
type MultiRewardPeriods []MultiRewardPeriod

// GetMultiRewardPeriod fetches a MultiRewardPeriod from an array of MultiRewardPeriods by its denom
func (mrps MultiRewardPeriods) GetMultiRewardPeriod(denom string) (MultiRewardPeriod, bool) {
	for _, rp := range mrps {
		if rp.CollateralType == denom {
			return rp, true
		}
	}
	return MultiRewardPeriod{}, false
}

// GetMultiRewardPeriodIndex returns the index of a MultiRewardPeriod inside array MultiRewardPeriods
func (mrps MultiRewardPeriods) GetMultiRewardPeriodIndex(denom string) (int, bool) {
	for i, rp := range mrps {
		if rp.CollateralType == denom {
			return i, true
		}
	}
	return -1, false
}

// Validate checks if all the RewardPeriods are valid and there are no duplicated
// entries.
func (mrps MultiRewardPeriods) Validate() error {
	seenPeriods := make(map[string]bool)
	for _, rp := range mrps {
		if seenPeriods[rp.CollateralType] {
			return fmt.Errorf("duplicated reward period with collateral type %s", rp.CollateralType)
		}

		if err := rp.Validate(); err != nil {
			return err
		}
		seenPeriods[rp.CollateralType] = true
	}

	return nil
}

// ---------------------- Reward indexes are used internally in the store ----------------------

// NewRewardIndex returns a new RewardIndex
func NewRewardIndex(collateralType string, factor sdk.Dec) RewardIndex {
	return RewardIndex{
		CollateralType: collateralType,
		RewardFactor:   factor,
	}
}

// Validate validates reward index
func (ri RewardIndex) Validate() error {
	if ri.RewardFactor.IsNegative() {
		return fmt.Errorf("reward factor value should be positive, is %s for %s", ri.RewardFactor, ri.CollateralType)
	}
	if strings.TrimSpace(ri.CollateralType) == "" {
		return fmt.Errorf("collateral type should not be empty")
	}
	return nil
}

// RewardIndexes slice of RewardIndex
type RewardIndexes []RewardIndex

// GetRewardIndex fetches a RewardIndex by its denom
func (ris RewardIndexes) GetRewardIndex(denom string) (RewardIndex, bool) {
	for _, ri := range ris {
		if ri.CollateralType == denom {
			return ri, true
		}
	}
	return RewardIndex{}, false
}

// GetFactorIndex gets the index of a specific reward index inside the array by its index
func (ris RewardIndexes) GetFactorIndex(denom string) (int, bool) {
	for i, ri := range ris {
		if ri.CollateralType == denom {
			return i, true
		}
	}
	return -1, false
}

// Validate validation for reward indexes
func (ris RewardIndexes) Validate() error {
	for _, ri := range ris {
		if err := ri.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// MultiRewardIndex stores reward accumulation information on multiple reward types
type MultiRewardIndex struct {
	CollateralType string        `json:"collateral_type" yaml:"collateral_type"`
	RewardIndexes  RewardIndexes `json:"reward_indexes" yaml:"reward_indexes"`
}

// NewMultiRewardIndex returns a new MultiRewardIndex
func NewMultiRewardIndex(collateralType string, indexes RewardIndexes) MultiRewardIndex {
	return MultiRewardIndex{
		CollateralType: collateralType,
		RewardIndexes:  indexes,
	}
}

// GetFactorIndex gets the index of a specific reward index inside the array by its index
func (mri MultiRewardIndex) GetFactorIndex(denom string) (int, bool) {
	for i, ri := range mri.RewardIndexes {
		if ri.CollateralType == denom {
			return i, true
		}
	}
	return -1, false
}

func (mri MultiRewardIndex) String() string {
	return fmt.Sprintf(`Collateral Type: %s, Reward Indexes: %s`, mri.CollateralType, mri.RewardIndexes)
}

// Validate validates multi-reward index
func (mri MultiRewardIndex) Validate() error {
	for _, rf := range mri.RewardIndexes {
		if rf.RewardFactor.IsNegative() {
			return fmt.Errorf("reward index's factor value cannot be negative: %s", rf)
		}
	}
	if strings.TrimSpace(mri.CollateralType) == "" {
		return fmt.Errorf("collateral type should not be empty")
	}
	return nil
}

// MultiRewardIndexes slice of MultiRewardIndex
type MultiRewardIndexes []MultiRewardIndex

// GetRewardIndex fetches a RewardIndex from a MultiRewardIndex by its denom
func (mris MultiRewardIndexes) GetRewardIndex(denom string) (MultiRewardIndex, bool) {
	for _, ri := range mris {
		if ri.CollateralType == denom {
			return ri, true
		}
	}
	return MultiRewardIndex{}, false
}

// GetRewardIndexIndex fetches a specific reward index inside the array by its denom
func (mris MultiRewardIndexes) GetRewardIndexIndex(denom string) (int, bool) {
	for i, ri := range mris {
		if ri.CollateralType == denom {
			return i, true
		}
	}
	return -1, false
}

// Validate validation for reward indexes
func (mris MultiRewardIndexes) Validate() error {
	for _, mri := range mris {
		if err := mri.Validate(); err != nil {
			return err
		}
	}
	return nil
}
