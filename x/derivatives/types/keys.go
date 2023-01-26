package types

import (
	"encoding/binary"
	"strconv"

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
	KeyPrefixPoolDeposit = "pool_deposit"
	//
	KeyPrefixOpenedPosition        = "opened_position"
	KeyPrefixClosedPosition        = "closed_position"
	KeyPrefixPerpetualFutures      = "perpetual_futures"
	KeyPrefixPerpetualOptions      = "perpetual_options"
	KeyPrefixNetPositionAmount     = "net_position_amount"
	KeyPrefixLastPositionId        = "last_position_id"
	KeyPrefixAccumulatedFee        = "accumulated_fee"
	KeyPrefixPoolMarketCapSnapshot = "pool_market_cap_snapshot"
	KeyPrefixLPTokenSupplySnapshot = "lpt_supply_snapshot"
	KeyPrefixPositionMargin        = "position_margin"
	KeyPrefixImaginaryFundingRate  = "imaginary_funding_rate"
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

func AddressPoolDepositKeyPrefix(depositor sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixPoolDeposit), address.MustLengthPrefix(depositor)...)
}

func AddressAssetPoolDepositKeyPrefix(depositor sdk.AccAddress, denom string) []byte {
	return append(append([]byte(KeyPrefixPoolDeposit), address.MustLengthPrefix(depositor)...), []byte(denom)...)
}

func AssetKeyPrefix(denom string) []byte {
	return append([]byte(KeyPrefixDerivativesPoolAssets), []byte(denom)...)
}

func AssetDepositKeyPrefix(denom string) []byte {
	return append([]byte(KeyPrefixPoolDeposit), []byte(denom)...)
}

func AddressOpenedPositionKeyPrefix(sender sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixOpenedPosition), address.MustLengthPrefix(sender)...)
}

func AddressClosedPositionKeyPrefix(sender sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixClosedPosition), address.MustLengthPrefix(sender)...)
}

func AddressOpenedPositionWithIdKeyPrefix(sender sdk.AccAddress, posId string) []byte {
	return append(AddressOpenedPositionKeyPrefix(sender), []byte(posId)...)
}

func AddressClosedPositionWithIdKeyPrefix(sender sdk.AccAddress, posId string) []byte {
	return append(AddressClosedPositionKeyPrefix(sender), []byte(posId)...)
}

func DenomNetPositionPerpetualFuturesKeyPrefix(denom string) []byte {
	return append(append([]byte(KeyPrefixPerpetualFutures), []byte(KeyPrefixNetPositionAmount)...), []byte(denom)...)
}

func AddressPoolMarketCapSnapshotKeyPrefix(height int64) []byte {
	return append([]byte(KeyPrefixPoolMarketCapSnapshot), []byte(strconv.FormatInt(height, 10))...)
}

func AddressLPTokenSupplySnapshotKeyPrefix(height int64) []byte {
	return append([]byte(KeyPrefixLPTokenSupplySnapshot), []byte(strconv.FormatInt(height, 10))...)
}

func RemainingMarginKeyPrefix(posId string) []byte {
	return append([]byte(KeyPrefixPositionMargin), []byte(posId)...)
}
