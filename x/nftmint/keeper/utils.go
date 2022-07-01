package keeper

import "encoding/binary"

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
