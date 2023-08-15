package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/cosmos/cosmos-sdk/x/nft/keeper"

	// wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/UnUniFi/chain/x/nft/types"
)

type Keeper struct {
	keeper.Keeper
	cdc codec.BinaryCodec

	// wasmKeeper wasmtypes.ContractOpsKeeper
	authority string
}

func NewKeeper(k keeper.Keeper, cdc codec.BinaryCodec /*wasmKeeper wasmtypes.ContractOpsKeeper,*/, authority string) Keeper {
	return Keeper{
		Keeper: k,
		cdc:    cdc,
		// wasmKeeper: wasmKeeper,
		authority: authority,
	}
}

func (k Keeper) GetClassData(ctx sdk.Context, classId string) (types.ClassData, bool) {
	token, found := k.Keeper.GetClass(ctx, classId)
	if !found {
		return types.ClassData{}, false
	}

	var classDataI types.ClassDataI
	if err := k.cdc.UnpackAny(token.Data, classDataI); err != nil {
		return types.ClassData{}, true
	}

	if token.Data == nil {
		return types.ClassData{}, true
	}

	switch classData := classDataI.(type) {
	case *types.ClassData:
		return *classData, true
	default:
		return types.ClassData{}, true
	}
}

func (k Keeper) SetClassData(ctx sdk.Context, classId string, data types.ClassData) error {
	class, found := k.Keeper.GetClass(ctx, classId)
	if !found {
		return nft.ErrClassNotExists
	}

	dataAny, err := codectypes.NewAnyWithValue(&data)
	if err != nil {
		return err
	}

	class.Data = dataAny
	err = k.Keeper.UpdateClass(ctx, class)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) GetNftData(ctx sdk.Context, classId string, id string) (types.NftData, bool) {
	token, found := k.Keeper.GetNFT(ctx, classId, id)
	if !found {
		return types.NftData{}, false
	}

	var nftDataI types.NftDataI
	if err := k.cdc.UnpackAny(token.Data, nftDataI); err != nil {
		return types.NftData{}, true
	}

	if token.Data == nil {
		return types.NftData{}, true
	}

	switch nftData := nftDataI.(type) {
	case *types.NftData:
		return *nftData, true
	default:
		return types.NftData{}, true
	}
}

func (k Keeper) SetNftData(ctx sdk.Context, classId string, id string, data types.NftData) error {
	token, found := k.Keeper.GetNFT(ctx, classId, id)
	if !found {
		return nft.ErrNFTNotExists
	}

	dataAny, err := codectypes.NewAnyWithValue(&data)
	if err != nil {
		return err
	}

	token.Data = dataAny
	err = k.Keeper.Update(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
