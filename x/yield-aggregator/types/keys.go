package types

import (
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "yield_aggregator"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	VaultKey         = "Vault/value/"
	VaultCountKey    = "Vault/count/"
	StrategyKey      = "Strategy/value/"
	StrategyCountKey = "Strategy/count/"
)

func KeyPrefixStrategy(vaultDenom string) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s/%s", VaultKey, vaultDenom, StrategyKey))
}

func KeyPrefixStrategyCount(vaultDenom string) []byte {
	return KeyPrefix(fmt.Sprintf("%s%s/%s", VaultKey, vaultDenom, StrategyCountKey))
}
