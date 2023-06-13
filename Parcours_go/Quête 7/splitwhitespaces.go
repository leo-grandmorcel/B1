package piscine

func SplitWhiteSpaces(s string) []string {
	result := []string{}
	var strTemp string
	for i := 0; i < len(s); i++ {
		if s[i] > 32 && s[i] <= 126 {
			strTemp += string(s[i])
		} else {
			if len(strTemp) > 0 {
				result = append(result, strTemp)
				strTemp = ""
			}
		}
	}
	result = append(result, strTemp)
	return result
}
