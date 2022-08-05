package keeper

import (
	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

var _ types.QueryServer = Keeper{}
