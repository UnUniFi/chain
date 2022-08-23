package types

const (
	// ModuleName defines the module name
	ModuleName = "decentralizedvault"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_decentralizedvault"
)

const (
	WrappedClassName         = "ununifi wrapped"
	WrappedClassId           = "ununifi-wrapped"
	WrappedClassSymbol       = "ununifi-wrapped"
	WrappedClassDescription  = "ununifi wrapped nft class"
	KeyPrefixTransferRequest = "transfer_req"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func TransferRequestKey(idBytes []byte) []byte {
	return append([]byte(KeyPrefixTransferRequest), idBytes...)
}
