package piscine

func StrRev(s string) string {
	long := []rune(s)
	for i, j := 0, len(long)-1; i < j; i, j = i+1, j-1 {
		long[i], long[j] = long[j], long[i]
	}
	return string(long)
}
