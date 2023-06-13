package piscine

func IsLower(s string) bool {
	result := false
	for _, letter := range s {
		if letter >= 97 && letter <= 122 {
			result = true
		} else {
			return false
		}
	}
	return result
}
