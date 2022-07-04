package keeper

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func SliceIndex(s []string, element string) int {
	for i := 0; i < len(s); i++ {
		if element == s[i] {
			return i
		}
	}
	return -1
}

func RemoveIndex(s []string, index int) []string {
	return (append(s[:index], s[index+1:]...))
}

func UintToByte(u uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, u)
	return b
}

// Create class id on UnUniFi using addr sequence and addr byte
func createClassId(num uint64, addr sdk.Address) string {
	sequenceByte := UintToByte(num)
	addrByte := addr.Bytes()
	idByte := append(addrByte, sequenceByte...)

	idHash := sha256.Sum256(idByte)
	idString := hex.EncodeToString(idHash[LenHashByteToHex:])
	classID := PrefixClassId + strings.ToUpper(idString)

	return classID
}
