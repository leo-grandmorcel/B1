package piscine

func FirstRune(s string) rune {
	for index, letter := range s {
		if index == 0 {
			return letter
		}
	}
	return 0
}
