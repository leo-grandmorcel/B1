package piscine

func BasicJoin(elems []string) string {
	result := ""
	for _, letter := range elems {
		result += string(letter)
	}
	return result
}
