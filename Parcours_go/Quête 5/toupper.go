package piscine

func ToUpper(s string) string {
	word := []rune(s)
	var result string = ""
	for index := range s {
		if word[index] >= 97 && word[index] <= 122 {
			word[index] = word[index] - 32
			result = result + string(word[index])
		} else {
			result = result + string(word[index])
		}
	}
	return result
}
