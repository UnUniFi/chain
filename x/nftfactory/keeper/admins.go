package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gogo/protobuf/proto"

	"github.com/UnUniFi/chain/x/nftfactory/types"
)

// GetAuthorityMetadata returns the authority metadata for a specific denom
func (k Keeper) GetAuthorityMetadata(ctx sdk.Context, denom string) (types.ClassAuthorityMetadata, error) {
	bz := k.GetDenomPrefixStore(ctx, denom).Get([]byte(types.ClassAuthorityMetadataKey))

	metadata := types.ClassAuthorityMetadata{}
	err := proto.Unmarshal(bz, &metadata)
	if err != nil {
		return types.ClassAuthorityMetadata{}, err
	}
	return metadata, nil
}

// setAuthorityMetadata stores authority metadata for a specific denom
func (k Keeper) setAuthorityMetadata(ctx sdk.Context, denom string, metadata types.ClassAuthorityMetadata) error {
	err := metadata.Validate()
	if err != nil {
		return err
	}

	store := k.GetDenomPrefixStore(ctx, denom)

	bz, err := proto.Marshal(&metadata)
	if err != nil {
		return err
	}

	store.Set([]byte(types.ClassAuthorityMetadataKey), bz)
	return nil
}

func (k Keeper) setAdmin(ctx sdk.Context, denom, admin string) error {
	metadata, err := k.GetAuthorityMetadata(ctx, denom)
	if err != nil {
		return err
	}

	metadata.Admin = admin

	return k.setAuthorityMetadata(ctx, denom, metadata)
}
