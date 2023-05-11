package types

import (
	"bytes"
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "cdp"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"

	// LiquidatorMacc module account for liquidator
	LiquidatorMacc = "liquidator"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ParamsKey = "Params-value-"
	CdpKey    = "Cdp-value-"
	// CdpCountKey          = "Cdp-count-"
	CdpIDKey                  = "CdpID-value-"
	CdpIDOwnerIndex           = "CdpID-Owner-index-"
	CdpIDCollateralRatioIndex = "CdpID-CollateralRatio-index-"
	NextCdpID                 = "NextCdpID"
	DebtDenom                 = "DebtDenom"
	GovDenom                  = "GovDenom"
	DepositKey                = "Deposit-value-"
	PrincipalKey              = "Principal-value-"
	PricefeedStatusKey        = "PricefeedStatus-value-"
	PreviousAccrualTime       = "PreviousAccrualTime-"
	InterestFactor            = "InterestFactor-"
)

var sep = []byte(":")

// GetCdpIDBytes returns the byte representation of the cdpID
func GetCdpIDBytes(cdpID uint64) (cdpIDBz []byte) {
	cdpIDBz = make([]byte, 8)
	binary.BigEndian.PutUint64(cdpIDBz, cdpID)
	return
}

// GetCdpIDFromBytes returns cdpID in uint64 format from a byte array
func GetCdpIDFromBytes(bz []byte) (cdpID uint64) {
	return binary.BigEndian.Uint64(bz)
}

// CdpKeySuffix key of a specific cdp in the store
func CdpKeySuffix(denomByte byte, cdpID uint64) []byte {
	return createKey([]byte{denomByte}, sep, GetCdpIDBytes(cdpID))
}

// SplitCdpKey returns the component parts of a cdp key
func SplitCdpKey(key []byte) (byte, uint64) {
	split := bytes.Split(key, sep)
	return split[0][0], GetCdpIDFromBytes(split[1])
}

// DenomIterKey returns the key for iterating over cdps of a certain denom in the store
func DenomIterKey(denomByte byte) []byte {
	return append([]byte{denomByte}, sep...)
}

// SplitDenomIterKey returns the component part of a key for iterating over cdps by denom
func SplitDenomIterKey(key []byte) byte {
	split := bytes.Split(key, sep)
	return split[0][0]
}

// DepositKeySuffix key of a specific deposit in the store
func DepositKeySuffix(cdpID uint64, depositor sdk.AccAddress) []byte {
	return createKey(GetCdpIDBytes(cdpID), sep, depositor)
}

// SplitDepositKey returns the component parts of a deposit key
func SplitDepositKey(key []byte) (uint64, sdk.AccAddress) {
	cdpID := GetCdpIDFromBytes(key[0:8])
	addr := key[9:]
	return cdpID, addr
}

// DepositIterKey returns the prefix key for iterating over deposits to a cdp
func DepositIterKey(cdpID uint64) []byte {
	return GetCdpIDBytes(cdpID)
}

// SplitDepositIterKey returns the component parts of a key for iterating over deposits on a cdp
func SplitDepositIterKey(key []byte) (cdpID uint64) {
	return GetCdpIDFromBytes(key)
}

// CollateralRatioBytes returns the liquidation ratio as sortable bytes
func CollateralRatioBytes(ratio sdk.Dec) []byte {
	ok := ValidSortableDec(ratio)
	if !ok {
		// set to max sortable if input is too large.
		ratio = sdk.OneDec().Quo(sdk.SmallestDec())
	}
	return SortableDecBytes(ratio)
}

// CollateralRatioKey returns the key for querying a cdp by its liquidation ratio
func CollateralRatioKey(denomByte byte, cdpID uint64, ratio sdk.Dec) []byte {
	ratioBytes := CollateralRatioBytes(ratio)
	idBytes := GetCdpIDBytes(cdpID)

	return createKey([]byte{denomByte}, sep, ratioBytes, sep, idBytes)
}

// SplitCollateralRatioKey split the collateral ratio key and return the denom, cdp id, and collateral:debt ratio
func SplitCollateralRatioKey(key []byte) (denom byte, cdpID uint64, ratio sdk.Dec) {

	cdpID = GetCdpIDFromBytes(key[len(key)-8:])
	split := bytes.Split(key[:len(key)-8], sep)
	denom = split[0][0]

	ratio, err := ParseDecBytes(split[1])
	if err != nil {
		panic(err)
	}
	return
}

// CollateralRatioIterKey returns the key for iterating over cdps by denom and liquidation ratio
func CollateralRatioIterKey(denomByte byte, ratio sdk.Dec) []byte {
	ratioBytes := CollateralRatioBytes(ratio)
	return createKey([]byte{denomByte}, sep, ratioBytes)
}

// SplitCollateralRatioIterKey split the collateral ratio key and return the denom, cdp id, and collateral:debt ratio
func SplitCollateralRatioIterKey(key []byte) (denom byte, ratio sdk.Dec) {
	split := bytes.Split(key, sep)
	denom = split[0][0]

	ratio, err := ParseDecBytes(split[1])
	if err != nil {
		panic(err)
	}
	return
}

func createKey(bytes ...[]byte) (r []byte) {
	for _, b := range bytes {
		r = append(r, b...)
	}
	return
}
