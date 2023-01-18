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
	KeyPrefixDerivativesPoolAssets = "pool_assets"
	// subpool assets
	KeyPrefixDerivativesSubpoolAssets = "subpool_assets"
	// user deposited real assets
	KeyPrefixDerivativesUserDepositedAssets = "user_deposited_assets"
	// User deposits by address
	KeyPrefixAddressDeposit = "address_deposit"
	// assets deposits by denom
	KeyPrefixAssetDeposit = "asset_deposit"
	//
	KeyPrefixPosition         = "position"
	KeyPrefixClosedPosition   = "closed_position"
	KeyPrefixPerpetualFutures = "perpetual_futures"
	KeyPrefixPerpetualOptions = "perpetual_options"
	KeyPrefixNetPosition      = "net_position"
)

func AddressDepositKeyPrefix(depositor sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixAddressDeposit), address.MustLengthPrefix(depositor)...)
}

func AddressAssetDepositKeyPrefix(depositor sdk.AccAddress, denom string) []byte {
	return append(append([]byte(KeyPrefixAddressDeposit), address.MustLengthPrefix(depositor)...), []byte(denom)...)
}

func AssetKeyPrefix(denom string) []byte {
	return append([]byte(KeyPrefixDerivativesPoolAssets), []byte(denom)...)
}

func AssetDepositKeyPrefix(denom string) []byte {
	return append([]byte(KeyPrefixAssetDeposit), []byte(denom)...)
}

func AddressPositionKeyPrefix(sender sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixPosition), address.MustLengthPrefix(sender)...)
}

func AddressPositionWithIdKeyPrefix(sender sdk.AccAddress, posId int) []byte {
	return append(AddressPositionKeyPrefix(sender), byte(posId))
}

func AddressClosedPositionKeyPrefix(sender sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixPosition), address.MustLengthPrefix(sender)...)
}

func AddressClosedPositionWithIdKeyPrefix(sender sdk.AccAddress, posId int) []byte {
	return append(AddressPositionKeyPrefix(sender), byte(posId))
}

func DenomNetPositionPerpetualFuturesKeyPrefix(denom string) []byte {
	return append(append([]byte(KeyPrefixPerpetualFutures), []byte(KeyPrefixNetPosition)...), []byte(denom)...)
}
