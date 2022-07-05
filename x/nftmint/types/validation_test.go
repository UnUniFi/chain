package types_test

import (
	"testing"

	"github.com/UnUniFi/chain/x/nftmint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	// For the test, use "cosmos" prefix
	testAddr  = "cosmos1nyd8wdqyrnjfwfnfysv6t0rrpcj4pmzkykytjh"
	testAddr2 = "cosmos1chjjqrherp2lgmj9wsqwe6leercydncqx2v209"
)

func TestValidateMintingPermission(t *testing.T) {
	// OnlyOwner case
	owner, _ := sdk.AccAddressFromBech32(testAddr)
	classAttirbutes := types.ClassAttributes{
		Owner:             owner.Bytes(),
		MintingPermission: 0,
	}
	err := types.ValidateMintingPermission(classAttirbutes, owner)
	require.NoError(t, err)

	falseCase, _ := sdk.AccAddressFromBech32(testAddr2)
	err = types.ValidateMintingPermission(classAttirbutes, falseCase)
	require.Error(t, err)

	// AnyOne case
	classAttirbutes = types.ClassAttributes{
		MintingPermission: 1,
	}
	err = types.ValidateMintingPermission(classAttirbutes, owner)
	require.NoError(t, err)

	// In case of now allowed option
	classAttirbutes = types.ClassAttributes{
		MintingPermission: 3,
	}
	err = types.ValidateMintingPermission(classAttirbutes, owner)
	require.Error(t, err)
}
