package piscine

func IsUpper(s string) bool {
	result := false
	for _, letter := range s {
		if letter >= 65 && letter <= 90 {
			result = true
		} else {
			result = false
		}
	}
	return result
}
