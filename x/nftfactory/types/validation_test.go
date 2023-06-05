package types_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/UnUniFi/chain/x/nftfactory/types"
)

const (
	// For the test, use "cosmos" prefix
	testAddr               = "cosmos1nyd8wdqyrnjfwfnfysv6t0rrpcj4pmzkykytjh"
	testAddr2              = "cosmos1chjjqrherp2lgmj9wsqwe6leercydncqx2v209"
	testAddr3              = "cosmos1hzgsxn26pn7zt0g3yssld7cmj86zr7fhhu2r3g"
	testClassName          = "UnUniFi"
	testUri                = "ipfs://test/"
	testTokenSupplyCap     = 10000
	testCurrentTokenSupply = 10
	testSymbol             = "TEST"
	testDescription        = "This description is for the valdation uni test"
)

func TestValidateMintingPermission(t *testing.T) {
	// OnlyOwner case
	owner, _ := sdk.AccAddressFromBech32(testAddr)
	classAttirbutes := types.ClassAttributes{
		Owner:             owner.Bytes(),
		MintingPermission: 0,
	}
	err := types.ValidateMintingPermission(classAttirbutes.MintingPermission, owner, owner)
	require.NoError(t, err)

	falseCase, _ := sdk.AccAddressFromBech32(testAddr2)
	err = types.ValidateMintingPermission(classAttirbutes.MintingPermission, owner, falseCase)
	require.Error(t, err)

	// AnyOne case
	classAttirbutes = types.ClassAttributes{
		MintingPermission: 1,
	}
	anyoneCase, _ := sdk.AccAddressFromBech32(testAddr3)
	err = types.ValidateMintingPermission(classAttirbutes.MintingPermission, owner, anyoneCase)
	require.NoError(t, err)

	// In case of now allowed option
	classAttirbutes = types.ClassAttributes{
		MintingPermission: 3,
	}
	err = types.ValidateMintingPermission(classAttirbutes.MintingPermission, nil, nil)
	require.Error(t, err)
}

func TestValidateClassName(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateClassName(params.MinClassNameLen, params.MaxClassNameLen, testClassName)
	require.NoError(t, err)

	// invalid case which name is shorter than the default MinClassNameLen
	invalidClassNameShort := testClassName[:2]
	err = types.ValidateClassName(params.MinClassNameLen, 10000, invalidClassNameShort)
	require.Error(t, err)

	// invalid case which name is longer than the default MaxClassNameLen
	var invalidClassNameLong string
	for i := 0; i <= (int(params.MaxClassNameLen) / 7); i++ {
		invalidClassNameLong += testClassName
	}
	err = types.ValidateClassName(0, params.MaxClassNameLen, invalidClassNameLong)
	require.Error(t, err)
}

func TestValidateUri(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateUri(params.MinUriLen, params.MaxUriLen, testUri)
	require.NoError(t, err)

	// invalid case which uri is shoter than the default MinUriLen
	invalidUriShort := testUri[:4]
	err = types.ValidateUri(params.MinUriLen, 10000, invalidUriShort)
	require.Error(t, err)

	// invalid case which uri is longer than the default MaxUriLen
	var invalidUriLong string
	for i := 0; i <= (int(params.MaxUriLen) / len(testUri)); i++ {
		invalidUriLong += testUri
	}
	err = types.ValidateUri(0, params.MaxUriLen, invalidUriLong)
	require.Error(t, err)
}

func TestValidateTokenSupplyCap(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateTokenSupplyCap(params.MaxNFTSupplyCap, testTokenSupplyCap)
	require.NoError(t, err)

	// invalid case which token supply cap is bigger than the default MaxTokenSupplyCap
	invalidTokenSupply := testTokenSupplyCap * ((params.MaxNFTSupplyCap)/testTokenSupplyCap + 1)
	err = types.ValidateTokenSupplyCap(params.MaxNFTSupplyCap, invalidTokenSupply)
	require.Error(t, err)
}

func TestValidateTokenSupply(t *testing.T) {
	err := types.ValidateTokenSupply(testCurrentTokenSupply, testTokenSupplyCap)
	require.NoError(t, err)

	// invalid case which current token supply is over the token supply cap
	invalidTokenSupplyCap := 5
	err = types.ValidateTokenSupply(testCurrentTokenSupply, uint64(invalidTokenSupplyCap))
	require.Error(t, err)
}

func TestValidateSymbol(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateSymbol(params.MaxSymbolLen, testSymbol)
	require.NoError(t, err)

	// invalid case which symbol is longer that the default MaxSymbolLen
	var invalidSymbol string
	for i := 0; i <= (int(params.MaxSymbolLen) / len(testSymbol)); i++ {
		invalidSymbol += testSymbol
	}
	err = types.ValidateSymbol(params.MaxSymbolLen, invalidSymbol)
	require.Error(t, err)
}

func TestValidateDescription(t *testing.T) {
	params := types.DefaultParams()

	// valid case
	err := types.ValidateDescription(params.MaxDescriptionLen, testDescription)
	require.NoError(t, err)

	// invalid case which description is longer than the default MaxDescriptionLen
	var invalidDescription string
	for i := 0; i <= (int(params.MaxDescriptionLen) / len(testDescription)); i++ {
		invalidDescription += testDescription
	}
	err = types.ValidateDescription(params.MaxDescriptionLen, invalidDescription)
	require.Error(t, err)
}
