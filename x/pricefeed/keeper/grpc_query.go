package keeper

import (
	"github.com/UnUniFi/chain/x/pricefeed/types"
)

var _ types.QueryServer = Keeper{}
