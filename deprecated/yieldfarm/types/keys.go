package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleName defines the module name
	ModuleName = "yieldfarm"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName
)

const (
	PrefixKeyFarmerInfo = "farmer_info_"
)

func FarmerInfoKey(addr sdk.AccAddress) []byte {
	return append([]byte(PrefixKeyFarmerInfo), addr...)
}
