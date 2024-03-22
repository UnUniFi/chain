package keeper

import (
	"github.com/UnUniFi/chain/x/irs/types"
)

var _ types.QueryServer = Keeper{}
