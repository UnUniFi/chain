package keeper

import (
	"github.com/UnUniFi/chain/x/auction/types"
)

var _ types.QueryServer = Keeper{}
