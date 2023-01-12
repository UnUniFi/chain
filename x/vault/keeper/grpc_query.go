package keeper

import (
	"github.com/UnUniFi/chain/x/vault/types"
)

var _ types.QueryServer = Keeper{}
