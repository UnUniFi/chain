package types

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/proto"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultNextAuctionID is the starting poiint for auction IDs.
const DefaultNextAuctionID uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	params := DefaultParams()
	return &GenesisState{
		NextAuctionId: DefaultNextAuctionID,
		Params:        params,
		Auctions:      []*types.Any{},
		// this line is used by starport scaffolding # genesis/types/default
		// AuctionList: []*Auction{},
	}
}

// NewGenesisState returns a new genesis state object for auctions module.
func NewGenesisState(nextID uint64, ap Params, ga []*types.Any) GenesisState {
	return GenesisState{
		NextAuctionId: nextID,
		Params:        ap,
		Auctions:      ga,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # genesis/types/validate
	// Check for duplicated ID in auction
	// auctionIdMap := make(map[string]bool)

	// for _, elem := range gs.AuctionList {
	// 	if _, ok := auctionIdMap[elem.Id]; ok {
	// 		return fmt.Errorf("duplicated id for auction")
	// 	}
	// 	auctionIdMap[elem.Id] = true
	// }
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	ids := map[uint64]bool{}
	auctions, err := UnpackGenesisAuctions(gs.Auctions)
	if err != nil {
		return err
	}
	for _, a := range auctions {

		if err := a.Validate(); err != nil {
			return fmt.Errorf("found invalid auction: %w", err)
		}

		if ids[a.GetID()] {
			return fmt.Errorf("found duplicate auction ID (%d)", a.GetID())
		}
		ids[a.GetID()] = true

		if a.GetID() >= gs.NextAuctionId {
			return fmt.Errorf("found auction ID ≥ the nextAuctionID (%d ≥ %d)", a.GetID(), gs.NextAuctionId)
		}
	}
	return nil
}

// Equal checks whether two GenesisState structs are equivalent.
func (gs GenesisState) Equal(gs2 GenesisState) bool {
	b1 := ModuleCdc.MustMarshalBinaryBare(&gs)
	b2 := ModuleCdc.MustMarshalBinaryBare(&gs2)
	return bytes.Equal(b1, b2)
}

// IsEmpty returns true if a GenesisState is empty.
func (gs GenesisState) IsEmpty() bool {
	return gs.Equal(GenesisState{})
}

// GenesisAuction is an interface that extends the auction interface to add functionality needed for initializing auctions from genesis.
type GenesisAuction interface {
	Auction
	GetModuleAccountCoins() sdk.Coins
	Validate() error
}

// GenesisAuctions is a slice of genesis auctions.
type GenesisAuctions []GenesisAuction

func PackGenesisAuctions(auctions GenesisAuctions) ([]*codectypes.Any, error) {
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

func UnpackGenesisAuctions(auctionsAny []*types.Any) (GenesisAuctions, error) {
	accounts := make(GenesisAuctions, len(auctionsAny))
	for i, any := range auctionsAny {
		acc, ok := any.GetCachedValue().(GenesisAuction)
		if !ok {
			return nil, fmt.Errorf("expected genesis account")
		}
		accounts[i] = acc
	}

	return accounts, nil
}
