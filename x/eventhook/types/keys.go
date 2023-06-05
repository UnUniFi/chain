package types

import (
	"fmt"
)

const (
	// ModuleName defines the module name
	ModuleName = "eventhook"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_eventhook"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	HookKey      = "Hook/value/"
	HookCountKey = "Hook/count/"
)

func KeyPrefixHook(eventType string) []byte {
	return KeyPrefix(fmt.Sprintf("%s/%s", HookKey, eventType))
}

func KeyPrefixHookCount(eventType string) []byte {
	return KeyPrefix(fmt.Sprintf("%s/%s", HookCountKey, eventType))
}
