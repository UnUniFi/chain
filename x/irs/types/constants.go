package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// OneShare represents the amount of subshares in a single pool share.
	OneShare = sdk.NewIntFromBigInt(sdk.OneDec().BigInt())
)
