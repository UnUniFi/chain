package types

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	proto "github.com/gogo/protobuf/proto"
	jpyx "github.com/lcnem/jpyx/types"
)

const (
	CollateralAuctionType = "collateral"
	SurplusAuctionType    = "surplus"
	DebtAuctionType       = "debt"
	ForwardAuctionPhase   = "forward"
	ReverseAuctionPhase   = "reverse"
)

// DistantFuture is a very large time value to use as initial the ending time for auctions.
// It is not set to the max time supported. This can cause problems with time comparisons, see https://stackoverflow.com/a/32620397.
// Also amino panics when encoding times ≥ the start of year 10000.
var DistantFuture = time.Date(9000, 1, 1, 0, 0, 0, 0, time.UTC)

// Auction is an interface for handling common actions on auctions.
type Auction interface {
	proto.Message

	GetID() uint64
	WithID(uint64) Auction

	GetInitiator() string
	GetLot() sdk.Coin
	GetBidder() sdk.AccAddress
	GetBid() sdk.Coin
	GetEndTime() time.Time

	GetType() string
	GetPhase() string
}

// Auctions is a slice of auctions.
type Auctions []Auction

func PackAuctions(auctions Auctions) ([]*codectypes.Any, error) {
	auctionAny := make([]*types.Any, len(auctions))
	for i, auc := range auctions {
		msg, ok := auc.(proto.Message)
		if !ok {
			return nil, fmt.Errorf("cannot proto marshal %T", auc)
		}
		any, err := types.NewAnyWithValue(msg)
		if err != nil {
			return nil, err
		}
		auctionAny[i] = any
	}

	return auctionAny, nil
}

func UnpackAuctions(auctionsAny []*types.Any) (Auctions, error) {
	accounts := make(Auctions, len(auctionsAny))
	for i, any := range auctionsAny {
		acc, ok := any.GetCachedValue().(Auction)
		if !ok {
			return nil, fmt.Errorf("expected genesis account")
		}
		accounts[i] = acc
	}

	return accounts, nil
}

// GetID is a getter for auction ID.
func (a BaseAuction) GetID() uint64 { return a.Id }

// GetBidder is a getter for auction Bidder.
func (a BaseAuction) GetBidder() sdk.AccAddress { return a.Bidder.AccAddress() }

// GetType returns the auction type. Used to identify auctions in event attributes.
func (a BaseAuction) GetType() string { return "base" }

// Validate verifies that the auction end time is before max end time
func (a BaseAuction) Validate() error {
	// ID can be 0 for surplus, debt and collateral auctions
	if strings.TrimSpace(a.Initiator) == "" {
		return errors.New("auction initiator cannot be blank")
	}
	if !a.Lot.IsValid() {
		return fmt.Errorf("invalid lot: %s", a.Lot)
	}
	// NOTE: bidder can be empty for Surplus and Collateral auctions
	if !a.Bidder.AccAddress().Empty() && len(a.Bidder.AccAddress()) != sdk.AddrLen {
		return fmt.Errorf("the expected bidder address length is %d, actual length is %d", sdk.AddrLen, len(a.Bidder.AccAddress()))
	}
	if !a.Bid.IsValid() {
		return fmt.Errorf("invalid bid: %s", a.Bid)
	}
	if a.EndTime.Unix() <= 0 || a.MaxEndTime.Unix() <= 0 {
		return errors.New("end time cannot be zero")
	}
	if a.EndTime.After(a.MaxEndTime) {
		return fmt.Errorf("MaxEndTime < EndTime (%s < %s)", a.MaxEndTime, a.EndTime)
	}
	return nil
}

// WithID returns an auction with the ID set.
func (a SurplusAuction) WithID(id uint64) Auction { a.Id = id; return &a }

// GetType returns the auction type. Used to identify auctions in event attributes.
func (a SurplusAuction) GetType() string { return SurplusAuctionType }

// GetModuleAccountCoins returns the total number of coins held in the module account for this auction.
// It is used in genesis initialize the module account correctly.
func (a SurplusAuction) GetModuleAccountCoins() sdk.Coins {
	// a.Bid is paid out on bids, so is never stored in the module account
	return sdk.NewCoins(a.Lot)
}

// GetPhase returns the direction of a surplus auction, which never changes.
func (a SurplusAuction) GetPhase() string { return ForwardAuctionPhase }

// NewSurplusAuction returns a new surplus auction.
func NewSurplusAuction(seller string, lot sdk.Coin, bidDenom string, endTime time.Time) SurplusAuction {
	auction := SurplusAuction{BaseAuction{
		// no ID
		Initiator:       seller,
		Lot:             lot,
		Bidder:          nil,
		Bid:             sdk.NewInt64Coin(bidDenom, 0),
		HasReceivedBids: false, // new auctions don't have any bids
		EndTime:         endTime,
		MaxEndTime:      endTime,
	}}
	return auction
}

// WithID returns an auction with the ID set.
func (a DebtAuction) WithID(id uint64) Auction { a.Id = id; return &a }

// GetType returns the auction type. Used to identify auctions in event attributes.
func (a DebtAuction) GetType() string { return DebtAuctionType }

// GetModuleAccountCoins returns the total number of coins held in the module account for this auction.
// It is used in genesis initialize the module account correctly.
func (a DebtAuction) GetModuleAccountCoins() sdk.Coins {
	// a.Lot is minted at auction close, so is never stored in the module account
	// a.Bid is paid out on bids, so is never stored in the module account
	return sdk.NewCoins(a.CorrespondingDebt)
}

// GetPhase returns the direction of a debt auction, which never changes.
func (a DebtAuction) GetPhase() string { return ReverseAuctionPhase }

// Validate validates the DebtAuction fields values.
func (a DebtAuction) Validate() error {
	if !a.CorrespondingDebt.IsValid() {
		return fmt.Errorf("invalid corresponding debt: %s", a.CorrespondingDebt)
	}
	return a.BaseAuction.Validate()
}

// NewDebtAuction returns a new debt auction.
func NewDebtAuction(buyerModAccName string, bid sdk.Coin, initialLot sdk.Coin, endTime time.Time, debt sdk.Coin) DebtAuction {
	// Note: Bidder is set to the initiator's module account address instead of module name. (when the first bid is placed, it is paid out to the initiator)
	// Setting to the module account address bypasses calling supply.SendCoinsFromModuleToModule, instead calls SendCoinsFromModuleToAccount.
	// This isn't a problem currently, but if additional logic/validation was added for sending to coins to Module Accounts, it would be bypassed.
	auction := DebtAuction{
		BaseAuction: BaseAuction{
			// no ID
			Initiator:       buyerModAccName,
			Lot:             initialLot,
			Bidder:          authtypes.NewModuleAddress(buyerModAccName).Bytes(), // send proceeds from the first bid to the buyer.
			Bid:             bid,                                                                // amount that the buyer is buying - doesn't change over course of auction
			HasReceivedBids: false,                                                              // new auctions don't have any bids
			EndTime:         endTime,
			MaxEndTime:      endTime,
		},
		CorrespondingDebt: debt,
	}
	return auction
}

// CollateralAuction is a two phase auction.
// Initially, in forward auction phase, bids can be placed up to a max bid.
// Then it switches to a reverse auction phase, where the initial amount up for auction is bid down.
// Unsold Lot is sent to LotReturns, being divided among the addresses by weight.
// Collateral auctions are normally used to sell off collateral seized from CDPs.

// WithID returns an auction with the ID set.
func (a CollateralAuction) WithID(id uint64) Auction { a.Id = id; return &a }

// GetType returns the auction type. Used to identify auctions in event attributes.
func (a CollateralAuction) GetType() string { return CollateralAuctionType }

// GetModuleAccountCoins returns the total number of coins held in the module account for this auction.
// It is used in genesis initialize the module account correctly.
func (a CollateralAuction) GetModuleAccountCoins() sdk.Coins {
	// a.Bid is paid out on bids, so is never stored in the module account
	return sdk.NewCoins(a.Lot).Add(sdk.NewCoins(a.CorrespondingDebt)...)
}

// IsReversePhase returns whether the auction has switched over to reverse phase or not.
// CollateralAuctions initially start in forward phase.
func (a CollateralAuction) IsReversePhase() bool {
	return a.Bid.IsEqual(a.MaxBid)
}

// GetPhase returns the direction of a collateral auction.
func (a CollateralAuction) GetPhase() string {
	if a.IsReversePhase() {
		return ReverseAuctionPhase
	}
	return ForwardAuctionPhase
}

// Validate validates the CollateralAuction fields values.
func (a CollateralAuction) Validate() error {
	if !a.CorrespondingDebt.IsValid() {
		return fmt.Errorf("invalid corresponding debt: %s", a.CorrespondingDebt)
	}
	if !a.MaxBid.IsValid() {
		return fmt.Errorf("invalid max bid: %s", a.MaxBid)
	}
	if err := a.LotReturns.Validate(); err != nil {
		return fmt.Errorf("invalid lot returns: %w", err)
	}
	return a.BaseAuction.Validate()
}

// NewCollateralAuction returns a new collateral auction.
func NewCollateralAuction(seller string, lot sdk.Coin, endTime time.Time, maxBid sdk.Coin, lotReturns WeightedAddresses, debt sdk.Coin) CollateralAuction {
	auction := CollateralAuction{
		BaseAuction: BaseAuction{
			// no ID
			Initiator:       seller,
			Lot:             lot,
			Bidder:          nil,
			Bid:             sdk.NewInt64Coin(maxBid.Denom, 0),
			HasReceivedBids: false, // new auctions don't have any bids
			EndTime:         endTime,
			MaxEndTime:      endTime},
		CorrespondingDebt: debt,
		MaxBid:            maxBid,
		LotReturns:        lotReturns,
	}
	return auction
}

// NewWeightedAddresses returns a new list addresses with weights.
func NewWeightedAddresses(addrs []sdk.AccAddress, weights []sdk.Int) (WeightedAddresses, error) {

	wa := WeightedAddresses{
		Addresses: jpyx.StringAccAddresses(addrs),
		Weights:   weights,
	}
	if err := wa.Validate(); err != nil {
		return WeightedAddresses{}, err
	}
	return wa, nil
}

// Validate checks for that the weights are not negative, not all zero, and the lengths match.
func (wa WeightedAddresses) Validate() error {
	if len(wa.Weights) < 1 {
		return fmt.Errorf("must be at least 1 weighted address")
	}

	if len(wa.Addresses) != len(wa.Weights) {
		return fmt.Errorf("number of addresses doesn't match number of weights, %d ≠ %d", len(wa.Addresses), len(wa.Weights))
	}

	totalWeight := sdk.ZeroInt()
	for i := range wa.Addresses {
		if wa.Addresses[i].AccAddress().Empty() {
			return fmt.Errorf("address %d cannot be empty", i)
		}
		if len(wa.Addresses[i]) != sdk.AddrLen {
			return fmt.Errorf("address %d has an invalid length: expected %d, got %d", i, sdk.AddrLen, len(wa.Addresses[i]))
		}
		if wa.Weights[i].IsNegative() {
			return fmt.Errorf("weight %d contains a negative amount: %s", i, wa.Weights[i])
		}
		totalWeight = totalWeight.Add(wa.Weights[i])
	}

	if !totalWeight.IsPositive() {
		return fmt.Errorf("total weight must be positive")
	}

	return nil
}
