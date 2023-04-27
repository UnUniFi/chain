package keeper

import (
	"github.com/UnUniFi/chain/x/deprecated/auction/types"
)

var _ types.QueryServer = Keeper{}
