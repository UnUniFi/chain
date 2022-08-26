package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/UnUniFi/chain/x/decentralized-vault/types"
)

type (
	Keeper struct {
		cdc             codec.Codec
		storeKey        storetypes.StoreKey
		memKey          storetypes.StoreKey
		paramSpace      paramtypes.Subspace
		accountKeeper   types.AccountKeeper
		nftKeeper       types.NftKeeper
		nftmarketKeeper types.NftMarketKeeper
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey,
	memKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace,
	accountKeeper types.AccountKeeper,
	nftKeeper types.NftKeeper,
	nftmarketKeeper types.NftMarketKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:             cdc,
		storeKey:        storeKey,
		memKey:          memKey,
		paramSpace:      paramSpace,
		accountKeeper:   accountKeeper,
		nftKeeper:       nftKeeper,
		nftmarketKeeper: nftmarketKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
