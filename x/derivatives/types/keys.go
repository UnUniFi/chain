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

	// MarginManager defines the margin manager for derivatives module
	MarginManager = "margin_manager"

	// position nft class id
	PositionNFTClassId = "derivatives/perpetual_futures/positions"

	// PendingPaymentManager defines the pending payment manager for derivatives module
	PendingPaymentManager = "pending_payment_manager"
)

const (
	// TODO: KeyPrefixDerivativesSubpoolAssets is unused. Remove it if it won't be necesary.
	// subpool assets
	KeyPrefixDerivativesSubpoolAssets = "subpool_assets"
	KeyPrefixPosition                 = "position"
	KeyPrefixUserPosition             = "user_position"
	KeyPrefixPendingPaymentPosition   = "pending_payment_position"
	KeyPrefixPerpetualFutures         = "perpetual_futures"
	KeyPrefixPerpetualOptions         = "perpetual_options"
	KeyPrefixGrossPositionAmount      = "gross_position_amount"
	KeyPrefixLastPositionId           = "last_position_id"
	// TODO: KeyPrefixAccumulatedFee is unused. Remove it if it won't be necesary.
	KeyPrefixAccumulatedFee        = "accumulated_fee"
	KeyPrefixPoolMarketCapSnapshot = "pool_market_cap_snapshot"
	KeyPrefixLPTokenSupplySnapshot = "lpt_supply_snapshot"
	KeyPrefixPositionMargin        = "position_margin"
	KeyPrefixImaginaryFundingRate  = "imaginary_funding_rate"
	KeyPrefixBlockTimestamp        = "block_timestamp"
	KeyPrefixLPTBaseRedeemFee      = "lpt_base_redeem_fee"
	KeyPrefixReservedCoin          = "reserved_coin"
)

const (
	LiquidityProviderTokenDenom = "udlp"
	OneMillionInt               = 1000000
	OneMillionString            = "1000000"
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

func GetPositionIdByteFromString(idStr string) []byte {
	intPosId, _ := strconv.Atoi(idStr)
	return GetPositionIdBytes(uint64(intPosId))
}

func GetBlockTimestampBytes(timestamp int64) (timestampBz []byte) {
	timestampBz = make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBz, uint64(timestamp))
	return
}

func GetBlockTimestampFromBytes(bz []byte) int64 {
	return int64(binary.BigEndian.Uint64(bz))
}

func PositionWithIdKeyPrefix(posId string) []byte {
	return append([]byte(KeyPrefixPosition), GetPositionIdByteFromString(posId)...)
}

func PendingPaymentPositionWithIdKeyPrefix(posId string) []byte {
	return append([]byte(KeyPrefixPendingPaymentPosition), GetPositionIdByteFromString(posId)...)
}

func AddressPositionKeyPrefix(sender sdk.AccAddress) []byte {
	return append([]byte(KeyPrefixUserPosition), address.MustLengthPrefix(sender)...)
}

func AddressPositionWithIdKeyPrefix(sender sdk.AccAddress, posId string) []byte {
	return append(AddressPositionKeyPrefix(sender), GetPositionIdByteFromString(posId)...)
}

func DenomGrossPositionPerpetualFuturesKeyPrefix(market Market, positionType PositionType) []byte {
	return append(append([]byte(KeyPrefixPerpetualFutures), []byte(KeyPrefixGrossPositionAmount)...), []byte(fmt.Sprintf("%s/%s/%s", market.BaseDenom, market.QuoteDenom, positionType))...)
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

func ReservedCoinKeyPrefix(marketType MarketType, denom string) []byte {
	return append([]byte(KeyPrefixReservedCoin), []byte(fmt.Sprintf("%s/%s", marketType, denom))...)
}
