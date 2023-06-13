package piscine

func Join(strs []string, sep string) string {
	var result string = ""
	for index, letter := range strs {
		result += string(letter)
		if index < len(strs)-1 {
			result += sep
		}
	}
	return result
}
