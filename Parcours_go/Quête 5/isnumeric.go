package piscine

func IsNumeric(s string) bool {
	result := false
	for _, letter := range s {
		if letter >= 48 && letter <= 57 {
			result = true
		} else {
			return false
		}
	}
	return result
}
