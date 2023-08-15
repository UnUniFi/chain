package types

import (
	proto "github.com/cosmos/gogoproto/proto"
)

type ClassDataI interface {
	proto.Message
}

type NftDataI interface {
	proto.Message
}
