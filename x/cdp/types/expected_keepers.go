package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// AccountKeeper expected interface for the account keeper (noalias)
type AccountKeeper interface {
}

// SupplyKeeper defines the expected supply keeper for module accounts  (noalias)
type BankKeeper interface {
}

// AuctionKeeper expected interface for the auction keeper (noalias)
type AuctionKeeper interface {
}

// PricefeedKeeper defines the expected interface for the pricefeed  (noalias)
type PricefeedKeeper interface {
}

// CDPHooks event hooks for other keepers to run code in response to CDP modifications
type CDPHooks interface {
	AfterCDPCreated(ctx sdk.Context, cdp CDP)
	BeforeCDPModified(ctx sdk.Context, cdp CDP)
}
