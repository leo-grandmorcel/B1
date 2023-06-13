package piscine

func NRune(s string, n int) rune {
	if len(s) < n || n <= 0 {
		return 0
	} else {
		word := []rune(s)
		return word[n-1]
	}
}
