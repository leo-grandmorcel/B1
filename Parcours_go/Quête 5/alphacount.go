package piscine

func AlphaCount(s string) int {
	word := []rune(s)
	cpt := 0
	for index := range s {
		if (word[index] >= 65 && word[index] <= 90) || (word[index] >= 97 && word[index] <= 122) {
			cpt++
		}
	}
	return cpt
}
