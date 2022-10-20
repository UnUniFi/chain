package keeper

import (
	"github.com/UnUniFi/chain/x/icacallbacks/types"
)

var _ types.QueryServer = Keeper{}
