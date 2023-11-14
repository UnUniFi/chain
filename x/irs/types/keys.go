package types

const (
	// ModuleName defines the module name
	ModuleName = "irs"

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
	VaultKey      = "Vault/value/"
	VaultCountKey = "Vault/count/"
)
