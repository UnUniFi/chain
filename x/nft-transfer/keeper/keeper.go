package keeper

import (
	nftkeeper "github.com/UnUniFi/chain/x/nft/keeper"
	"github.com/bianjieai/nft-transfer/keeper"
)

type Keeper struct {
	keeper.Keeper
	nftKeeper nftkeeper.Keeper
}

func NewKeeper(k keeper.Keeper, nftKeeper nftkeeper.Keeper) Keeper {
	return Keeper{
		Keeper:    k,
		nftKeeper: nftKeeper,
	}
}
