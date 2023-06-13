package piscine

func Split(s, sep string) []string {
	result := []string{}
	strTemp := ""
	for i := 0; i < len(s); i++ {
		if i < len(s)-len(sep) {
			if string(s[i:i+len(sep)]) == sep {
				result = append(result, strTemp)
				i += len(sep)
				strTemp = ""
			}
		}
		strTemp += string(s[i])
	}
	result = append(result, strTemp)
	return result
}
