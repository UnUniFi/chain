package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"

	tmtime "github.com/tendermint/tendermint/libs/time"

	cdptypes "github.com/UnUniFi/chain/x/cdp/types"
	ununifidistTypes "github.com/UnUniFi/chain/x/ununifidist/types"
)

// Valid reward multipliers
const (
	Small  MultiplierName = "small"
	Medium MultiplierName = "medium"
	Large  MultiplierName = "large"
)

// Parameter keys and default values
var (
	KeyCdpMintingRewardPeriods      = []byte("CdpMintingRewardPeriods")
	KeyClaimEnd                     = []byte("ClaimEnd")
	KeyMultipliers                  = []byte("ClaimMultipliers")
	DefaultActive                   = false
	DefaultRewardPeriods            = RewardPeriods{}
	DefaultMultiRewardPeriods       = MultiRewardPeriods{}
	DefaultMultipliers              = Multipliers{}
	DefaultCdpClaims                = CdpMintingClaims{}
	DefaultGenesisAccumulationTimes = GenesisAccumulationTimes{}
	DefaultClaimEnd                 = tmtime.Canonical(time.Unix(1, 0))
	DefaultPrincipalDenom           = cdptypes.DefaultStableDenom
	DefaultCDPMintingRewardDenom    = cdptypes.DefaultGovDenom
	IncentiveMacc                   = ununifidistTypes.ModuleName
)

// NewParams returns a new params object
func NewParams(cdpMinting RewardPeriods, multipliers Multipliers, claimEnd time.Time) Params {
	return Params{
		CdpMintingRewardPeriods: cdpMinting,
		ClaimMultipliers:        multipliers,
		ClaimEnd:                claimEnd,
	}
}

// DefaultParams returns default params for incentive module
func DefaultParams() Params {
	return NewParams(DefaultRewardPeriods, DefaultMultipliers, DefaultClaimEnd)
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{
		paramstype.NewParamSetPair(KeyCdpMintingRewardPeriods, &p.CdpMintingRewardPeriods, validateRewardPeriodsParam),
		paramstype.NewParamSetPair(KeyClaimEnd, &p.ClaimEnd, validateClaimEndParam),
		paramstype.NewParamSetPair(KeyMultipliers, &p.ClaimMultipliers, validateMultipliersParam),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {

	if err := validateMultipliersParam(p.ClaimMultipliers); err != nil {
		return err
	}

	if err := validateRewardPeriodsParam(p.CdpMintingRewardPeriods); err != nil {
		return err
	}

	return nil
}

func validateRewardPeriodsParam(i interface{}) error {
	rewards, ok := i.(RewardPeriods)
	if !ok {
		rewards, ok = i.([]RewardPeriod)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
	}

	return rewards.Validate()
}

func validateMultiRewardPeriodsParam(i interface{}) error {
	rewards, ok := i.(MultiRewardPeriods)
	if !ok {
		rewards, ok = i.([]MultiRewardPeriod)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
	}

	return rewards.Validate()
}

func validateMultipliersParam(i interface{}) error {
	multipliers, ok := i.(Multipliers)
	if !ok {
		multipliers, ok = i.([]Multiplier)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
	}
	return multipliers.Validate()
}

func validateClaimEndParam(i interface{}) error {
	endTime, ok := i.(time.Time)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if endTime.Unix() <= 0 {
		return fmt.Errorf("end time should not be zero")
	}
	return nil
}

// NewRewardPeriod returns a new RewardPeriod
func NewRewardPeriod(active bool, collateralType string, start time.Time, end time.Time, reward sdk.Coin) RewardPeriod {
	return RewardPeriod{
		Active:           active,
		CollateralType:   collateralType,
		Start:            start,
		End:              end,
		RewardsPerSecond: reward,
	}
}

// Validate performs a basic check of a RewardPeriod fields.
func (rp RewardPeriod) Validate() error {
	if rp.Start.Unix() <= 0 {
		return errors.New("reward period start time cannot be 0")
	}
	if rp.End.Unix() <= 0 {
		return errors.New("reward period end time cannot be 0")
	}
	if rp.Start.After(rp.End) {
		return fmt.Errorf("end period time %s cannot be before start time %s", rp.End, rp.Start)
	}
	if !rp.RewardsPerSecond.IsValid() {
		return fmt.Errorf("invalid reward amount: %s", rp.RewardsPerSecond)
	}
	if strings.TrimSpace(rp.CollateralType) == "" {
		return fmt.Errorf("reward period collateral type cannot be blank: %s", rp.String())
	}
	return nil
}

// RewardPeriods array of RewardPeriod
type RewardPeriods []RewardPeriod

// Validate checks if all the RewardPeriods are valid and there are no duplicated
// entries.
func (rps RewardPeriods) Validate() error {
	seenPeriods := make(map[string]bool)
	for _, rp := range rps {
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

// NewMultiplier returns a new Multiplier
func NewMultiplier(name MultiplierName, lockup int64, factor sdk.Dec) Multiplier {
	return Multiplier{
		Name:         string(name),
		MonthsLockup: lockup,
		Factor:       factor,
	}
}

// Validate multiplier param
func (m Multiplier) Validate() error {
	if err := MultiplierName(m.Name).IsValid(); err != nil {
		return err
	}
	if m.MonthsLockup < 0 {
		return fmt.Errorf("expected non-negative lockup, got %d", m.MonthsLockup)
	}
	if m.Factor.IsNegative() {
		return fmt.Errorf("expected non-negative factor, got %s", m.Factor.String())
	}

	return nil
}

// Multipliers slice of Multiplier
type Multipliers []Multiplier

// Validate validates each multiplier
func (ms Multipliers) Validate() error {
	for _, m := range ms {
		if err := m.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// String implements fmt.Stringer
func (ms Multipliers) String() string {
	out := "Claim Multipliers\n"
	for _, s := range ms {
		out += fmt.Sprintf("%s\n", s.String())
	}
	return out
}

// MultiplierName name for valid multiplier
type MultiplierName string

// IsValid checks if the input is one of the expected strings
func (mn MultiplierName) IsValid() error {
	switch mn {
	case Small, Medium, Large:
		return nil
	}
	return fmt.Errorf("invalid multiplier name: %s", mn)
}
