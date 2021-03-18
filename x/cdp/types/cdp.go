package types

import (
	"errors"
	fmt "fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewCDP creates a new CDP object
func NewCDP(id uint64, owner sdk.AccAddress, collateral sdk.Coin, collateralType string, principal sdk.Coin, time time.Time, interestFactor sdk.Dec) CDP {
	fees := sdk.NewCoin(principal.Denom, sdk.ZeroInt())
	return CDP{
		Id:              id,
		Owner:           owner.Bytes(),
		Type:            collateralType,
		Collateral:      collateral,
		Principal:       principal,
		AccumulatedFees: fees,
		FeesUpdated:     time,
		InterestFactor:  interestFactor,
	}
}

// NewCDPWithFees creates a new CDP object, for use during migration
func NewCDPWithFees(id uint64, owner sdk.AccAddress, collateral sdk.Coin, collateralType string, principal, fees sdk.Coin, time time.Time, interestFactor sdk.Dec) CDP {
	return CDP{
		Id:              id,
		Owner:           owner.Bytes(),
		Type:            collateralType,
		Collateral:      collateral,
		Principal:       principal,
		AccumulatedFees: fees,
		FeesUpdated:     time,
		InterestFactor:  interestFactor,
	}
}

// Validate performs a basic validation of the CDP fields.
func (cdp CDP) Validate() error {
	if cdp.Id == 0 {
		return errors.New("cdp id cannot be 0")
	}
	if cdp.Owner.AccAddress().Empty() {
		return errors.New("cdp owner cannot be empty")
	}
	if !cdp.Collateral.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "collateral %s", cdp.Collateral)
	}
	if !cdp.Principal.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "principal %s", cdp.Principal)
	}
	if !cdp.AccumulatedFees.IsValid() {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "accumulated fees %s", cdp.AccumulatedFees)
	}
	if cdp.FeesUpdated.Unix() <= 0 {
		return errors.New("cdp updated fee time cannot be zero")
	}
	if strings.TrimSpace(cdp.Type) == "" {
		return fmt.Errorf("cdp type cannot be empty")
	}
	return nil
}

// GetTotalPrincipal returns the total principle for the cdp
func (cdp CDP) GetTotalPrincipal() sdk.Coin {
	return cdp.Principal.Add(cdp.AccumulatedFees)
}

// CDPs a collection of CDP objects
type CDPs []CDP

// String implements stringer
func (cdps CDPs) String() string {
	out := ""
	for _, cdp := range cdps {
		out += cdp.String() + "\n"
	}
	return out
}

// Validate validates each CDP
func (cdps CDPs) Validate() error {
	for _, cdp := range cdps {
		if err := cdp.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// NewAugmentedCDP creates a new AugmentedCDP object
func NewAugmentedCDP(cdp CDP, collateralValue sdk.Coin, collateralizationRatio sdk.Dec) AugmentedCDP {
	augmentedCDP := AugmentedCDP{
		CDP: CDP{
			Id:              cdp.Id,
			Owner:           cdp.Owner,
			Type:            cdp.Type,
			Collateral:      cdp.Collateral,
			Principal:       cdp.Principal,
			AccumulatedFees: cdp.AccumulatedFees,
			FeesUpdated:     cdp.FeesUpdated,
			InterestFactor:  cdp.InterestFactor,
		},
		CollateralValue:        collateralValue,
		CollateralizationRatio: collateralizationRatio,
	}
	return augmentedCDP
}

// AugmentedCDPs a collection of AugmentedCDP objects
type AugmentedCDPs []AugmentedCDP

// String implements stringer
func (augcdps AugmentedCDPs) String() string {
	out := ""
	for _, augcdp := range augcdps {
		out += augcdp.String() + "\n"
	}
	return out
}
