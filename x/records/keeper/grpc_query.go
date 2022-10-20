package keeper

import (
	"github.com/UnUniFi/chain/x/records/types"
)

var _ types.QueryServer = Keeper{}
