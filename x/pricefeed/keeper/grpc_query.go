package keeper

import (
	"github.com/lcnem/jpyx/x/pricefeed/types"
)

var _ types.QueryServer = Keeper{}
