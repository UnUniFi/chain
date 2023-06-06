package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "yieldaggregatorv1"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PrefixKeyAssetManagementAccount = "asset_management_account_"
	PrefixKeyAssetManagementTarget  = "asset_management_target_"
	PrefixKeyFarmingOrder           = "farming_order_"
	PrefixKeyFarmingUnit            = "farming_unit_"
	PrefixKeyUserDeposit            = "user_deposit_"
	KeyLastFarmingUnit              = "last_farming_unit_"
	PrefixKeyDailyPercent           = "daily_percent_"
)

func AssetManagementAccountKey(id string) []byte {
	return append([]byte(PrefixKeyAssetManagementAccount), id...)
}

func AssetManagementTargetKey(accountId, targetId string) []byte {
	return append(append([]byte(PrefixKeyAssetManagementTarget), accountId...), targetId...)
}

func FarmingOrderKey(sender sdk.AccAddress, orderId string) []byte {
	return append(append([]byte(PrefixKeyFarmingOrder), sender...), orderId...)
}

func FarmingUnitKey(owner string, accId, targetId string) []byte {
	return append(append(append([]byte(PrefixKeyFarmingUnit), owner...), accId...), targetId...)
}

func DailyRewardKey(accId, targetId string) []byte {
	return append(append([]byte(PrefixKeyDailyPercent), accId...), targetId...)
}

func UserDepositKey(user sdk.AccAddress) []byte {
	return append([]byte(PrefixKeyUserDeposit), user...)
}
