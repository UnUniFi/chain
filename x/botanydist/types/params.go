package types

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstype "github.com/cosmos/cosmos-sdk/x/params/types"

	tmtime "github.com/tendermint/tendermint/types/time"

	cdptypes "github.com/lcnem/jpyx/x/cdp/types"
)

// Parameter keys and default values
var (
	KeyActive                = []byte("Active")
	KeyPeriods               = []byte("Periods")
	DefaultActive            = false
	DefaultPeriods           = Periods{}
	DefaultPreviousBlockTime = tmtime.Canonical(time.Unix(1, 0))
	DefaultGovDenom          = cdptypes.DefaultGovDenom
)

// NewPeriod returns a new instance of Period
func NewPeriod(start time.Time, end time.Time, inflation sdk.Dec) Period {
	return Period{
		Start:     start,
		End:       end,
		Inflation: inflation,
	}
}

// Periods array of Period
type Periods []Period

// String implements fmt.Stringer
func (prs Periods) String() string {
	out := "Periods\n"
	for _, pr := range prs {
		out += fmt.Sprintf("%s\n", pr)
	}
	return out
}

// NewParams returns a new params object
func NewParams(active bool, periods Periods) Params {
	return Params{
		Active:  active,
		Periods: periods,
	}
}

// DefaultParams returns default params for botanydist module
func DefaultParams() Params {
	return NewParams(DefaultActive, DefaultPeriods)
}

// ParamKeyTable Key declaration for parameters
func ParamKeyTable() paramstype.KeyTable {
	return paramstype.NewKeyTable().RegisterParamSet(&Params{})
}

// ParamSetPairs implements the ParamSet interface and returns all the key/value pairs
func (p *Params) ParamSetPairs() paramstype.ParamSetPairs {
	return paramstype.ParamSetPairs{
		paramstype.NewParamSetPair(KeyActive, &p.Active, validateActiveParam),
		paramstype.NewParamSetPair(KeyPeriods, &p.Periods, validatePeriodsParams),
	}
}

// Validate checks that the parameters have valid values.
func (p Params) Validate() error {
	if err := validateActiveParam(p.Active); err != nil {
		return err
	}

	return validatePeriodsParams(p.Periods)
}

func validateActiveParam(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}

func validatePeriodsParams(i interface{}) error {
	periods, ok := i.(Periods)
	if !ok {
		periods, ok = i.([]Period)
		if !ok {
			return fmt.Errorf("invalid parameter type: %T", i)
		}
	}

	prevEnd := tmtime.Canonical(time.Unix(0, 0))
	for _, pr := range periods {
		if pr.End.Before(pr.Start) {
			return fmt.Errorf("end time for period is before start time: %s", pr)
		}

		if pr.Start.Before(prevEnd) {
			return fmt.Errorf("periods must be in chronological order: %s", periods)
		}
		prevEnd = pr.End

		if pr.Start.Unix() <= 0 || pr.End.Unix() <= 0 {
			return fmt.Errorf("start or end time cannot be zero: %s", pr)
		}

		//TODO: validate period Inflation?
	}

	return nil
}
