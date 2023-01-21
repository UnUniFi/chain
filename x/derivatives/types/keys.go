package types

import (
	"encoding/binary"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"strconv"
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
	KeyPrefixPosition              = "position"
	KeyPrefixClosedPosition        = "closed_position"
	KeyPrefixPerpetualFutures      = "perpetual_futures"
	KeyPrefixPerpetualOptions      = "perpetual_options"
	KeyPrefixNetPosition           = "net_position"
	KeyPrefixLastPositionId        = "last_position_id"
	KeyPrefixAccumulatedFee        = "accumulated_fee"
	KeyPrefixPoolMarketCapSnapshot = "pool_market_cap_snapshot"
	KeyPrefixLPTokenSupplySnapshot = "lpt_supply_snapshot"
	KeyPrefixAPYMeasureId          = "apy_measure_id"
)

const (
	LiquidityProviderTokenDenom = "udlp"
)

func GetPositionIdBytes(posId uint64) (posIdBz []byte) {
	posIdBz = make([]byte, 8)
	binary.BigEndian.PutUint64(posIdBz, posId)
	return
}

func GetPositionIdFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}

func GetPositionIdFromString(idStr string) uint64 {
	return GetPositionIdFromBytes([]byte(idStr))
}

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

func AddressPositionWithIdKeyPrefix(sender sdk.AccAddress, posId uint64) []byte {
	return append(AddressPositionKeyPrefix(sender), GetPositionIdBytes(posId)...)
}

func AddressClosedPositionKeyPrefix(sender sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixPosition), address.MustLengthPrefix(sender)...)
}

func AddressClosedPositionWithIdKeyPrefix(sender sdk.AccAddress, posId uint64) []byte {
	return append(AddressPositionKeyPrefix(sender), GetPositionIdBytes(posId)...)
}

func DenomNetPositionPerpetualFuturesKeyPrefix(denom string) []byte {
	return append(append([]byte(KeyPrefixPerpetualFutures), []byte(KeyPrefixNetPosition)...), []byte(denom)...)
}

func AddressPoolMarketCapSnapshotKeyPrefix(height int64) []byte {
	return append([]byte(KeyPrefixPoolMarketCapSnapshot), []byte(strconv.FormatInt(height, 10))...)
}

func AddressLPTokenSupplySnapshotKeyPrefix(height int64) []byte {
	return append([]byte(KeyPrefixLPTokenSupplySnapshot), []byte(strconv.FormatInt(height, 10))...)
}
