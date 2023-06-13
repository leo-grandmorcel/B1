package piscine

func Decipher(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			if s[i]-16 < 65 {
				temp := s[i] - 64
				temp = 16 - temp
				result += string(rune(90 - temp))
			} else {
				result += string(rune(s[i] - 16))
			}
		} else if s[i] >= 97 && s[i] <= 122 {
			if s[i]-16 < 97 {
				temp := s[i] - 96
				temp = 16 - temp
				result += string(rune(122 - temp))
			} else {
				result += string(rune(s[i] - 16))
			}
		} else if s[i] == 32 {
			result += " "
		}
	}
	result += "\n"
	return result
}
