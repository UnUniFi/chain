package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/UnUniFi/chain/x/nftmint/types"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

type Keeper struct {
	cdc           codec.Codec
	storeKey      storetypes.StoreKey
	memKey        storetypes.StoreKey
	paramSpace    paramtypes.Subspace
	accountKeeper types.AccountKeeper
	nftKeeper     types.NftKeeper
}

func NewKeeper(cdc codec.Codec, storeKey, memKey storetypes.StoreKey,
	paramSpace paramtypes.Subspace, accountKeeper types.AccountKeeper,
	nftKeeper types.NftKeeper) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		paramSpace:    paramSpace,
		accountKeeper: accountKeeper,
		nftKeeper:     nftKeeper,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) CreateClass(
	ctx sdk.Context,
	classID, name, symbol, description, classUri string,
) error {
	return k.nftKeeper.SaveClass(ctx, types.NewClass(classID, name, symbol, description, classUri))
}

func (k Keeper) CreateClassAttributes(
	ctx sdk.Context,
	classID string,
	owner sdk.AccAddress,
	baseTokenUri string,
	mintingPermission types.MintingPermission,
	tokenSupplyCap uint64,
) {
	k.SaveClassAttributes(
		ctx,
		types.NewClassAttributes(
			classID,
			owner,
			baseTokenUri,
			mintingPermission,
			tokenSupplyCap,
		),
	)
}

func (k Keeper) MintNFT(
	ctx sdk.Context,
	classID, nftID string,
	owner sdk.AccAddress,
) error {
	classAttributes, exists := k.GetClassAttributes(ctx, classID)
	if !exists {
		return sdkerrors.Wrap(nfttypes.ErrClassNotExists, classID)
	}

	nftUri := classAttributes.BaseTokenUri + nftID
	return k.nftKeeper.Mint(
		ctx,
		types.NewNFT(
			classID,
			nftID,
			nftUri,
		),
		owner,
	)
}

func (k Keeper) CreateNFTAttributes(
	ctx sdk.Context,
	classID, nftID string,
	minter sdk.AccAddress,
) error {
	return k.SaveNFTAttributes(
		ctx,
		types.NewNFTAttributes(
			classID,
			nftID,
			minter,
		),
	)
}
