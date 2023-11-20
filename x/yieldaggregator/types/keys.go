package types

import (
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "yieldaggregator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	// Module parameters
	KeyParams = []byte{0x00}
)

const (
	VaultKey         = "Vault/value/"
	VaultCountKey    = "Vault/count/"
	StrategyKey      = "Strategy/value/"
	StrategyCountKey = "Strategy/count/"
	DenomInfoKey     = "Denom/info/"
	SymbolInfoKey    = "Symbol/info/"
	ChainReceiverKey = "ChainReceiver/info/"
)

func KeyPrefixStrategy(vaultDenom string) []byte {
	return KeyPrefix(fmt.Sprintf("%s/%s", StrategyKey, vaultDenom))
}

func KeyPrefixStrategyCount(vaultDenom string) []byte {
	return KeyPrefix(fmt.Sprintf("%s/%s", StrategyCountKey, vaultDenom))
}
