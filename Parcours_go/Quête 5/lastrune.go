package piscine

func LastRune(s string) rune {
	word := []rune(s)
	return word[len(s)-1]
}
