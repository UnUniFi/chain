package keeper

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
