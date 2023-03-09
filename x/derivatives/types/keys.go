package types

import (
	"encoding/binary"
	"fmt"
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

	// DerivativeFeeCollector defines the fee collector for derivatives module
	DerivativeFeeCollector = "derivatives_fee_collector"
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
	KeyPrefixPosition              = "position"
	KeyPrefixUserPosition          = "user_position"
	KeyPrefixPerpetualFutures      = "perpetual_futures"
	KeyPrefixPerpetualOptions      = "perpetual_options"
	KeyPrefixNetPositionAmount     = "net_position_amount"
	KeyPrefixLastPositionId        = "last_position_id"
	KeyPrefixAccumulatedFee        = "accumulated_fee"
	KeyPrefixPoolMarketCapSnapshot = "pool_market_cap_snapshot"
	KeyPrefixLPTokenSupplySnapshot = "lpt_supply_snapshot"
	KeyPrefixPositionMargin        = "position_margin"
	KeyPrefixImaginaryFundingRate  = "imaginary_funding_rate"
	KeyPrefixBlockTimestamp        = "block_timestamp"
	KeyPrefixLPTBaseRedeemFee      = "lpt_base_redeem_fee"
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

func GetBlockTimestampBytes(timestamp int64) (timestampBz []byte) {
	timestampBz = make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBz, uint64(timestamp))
	return
}

func GetBlockTimestampFromBytes(bz []byte) int64 {
	return int64(binary.BigEndian.Uint64(bz))
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

func PositionWithIdKeyPrefix(posId string) []byte {
	return append([]byte(KeyPrefixPosition), []byte(posId)...)
}

func AddressPositionKeyPrefix(sender sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixUserPosition), address.MustLengthPrefix(sender)...)
}

func AddressPositionWithIdKeyPrefix(sender sdk.AccAddress, posId string) []byte {
	return append(AddressPositionKeyPrefix(sender), []byte(posId)...)
}

func DenomNetPositionPerpetualFuturesKeyPrefix(denom string, quoteDenom string) []byte {
	return append(append([]byte(KeyPrefixPerpetualFutures), []byte(KeyPrefixNetPositionAmount)...), []byte(fmt.Sprintf("%s/%s", denom, quoteDenom))...)
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

func BlockTimestampWithHeight(height int64) []byte {
	return append([]byte(KeyPrefixBlockTimestamp), []byte(strconv.FormatInt(height, 10))...)
}
