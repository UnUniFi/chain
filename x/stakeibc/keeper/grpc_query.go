package keeper

import (
	"github.com/UnUniFi/chain/x/stakeibc/types"
)

var _ types.QueryServer = Keeper{}
