package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// OneShare represents the amount of subshares in a single pool share.
	OneShare = sdk.NewIntFromUint64(1_000_000)
)
