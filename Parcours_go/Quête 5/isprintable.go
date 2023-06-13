package piscine

func IsPrintable(s string) bool {
	result := false
	for _, letter := range s {
		if letter >= 32 && letter <= 126 {
			result = true
		} else {
			return false
		}
	}
	return result
}
