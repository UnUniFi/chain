package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (
	// ModuleName defines the module name
	ModuleName = "derivatives"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_derivatives"
)

const (
	// pool assets
	KeyPrefixDerivativesPoolAssets = "derivatives_pool_assets"
	// subpool assets
	KeyPrefixDerivativesSubpoolAssets = "derivatives_subpool_assets"
	// user deposited real assets
	KeyPrefixDerivativesUserDepositedAssets = "derivatives_user_deposited_assets"
	// User deposits by address
	KeyPrefixAddressDeposit = "address_deposit"
)

func AddressDepositKeyPrefix(depositor sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixAddressDeposit), address.MustLengthPrefix(depositor)...)
}

func AssetKeyPrefix(denom string) []byte {
	return append([]byte(KeyPrefixDerivativesPoolAssets), []byte(denom)...)
}
