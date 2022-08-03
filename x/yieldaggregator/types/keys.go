package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "yieldaggregator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_yieldaggregator"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	PrefixKeyAssetManagementAccount = "asset_management_account_"
	PrefixKeyAssetManagementTarget  = "asset_management_target_"
	PrefixKeyFarmingOrder           = "farming_order_"
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
