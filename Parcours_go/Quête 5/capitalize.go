package piscine

func Capitalize(s string) string {
	sC := []rune(s)
	cpt := 1
	result := ""
	for i := 0; i < len(s); i++ {
		if (sC[i] < 97 && sC[i] > 90) || (sC[i] < 48 && sC[i] > 0) || (sC[i] < 65 && sC[i] > 57) || sC[i] > 122 {
			cpt += 1
			result += string(sC[i])
		} else {
			if cpt > 0 {
				if sC[i] > 96 && sC[i] < 123 {
					sC[i] = sC[i] - 32
				}
				cpt, result = 0, result+string(sC[i])
			} else {
				if sC[i] > 64 && sC[i] < 91 {
					sC[i] = sC[i] + 32
				}
				result += string(sC[i])
			}
		}
	}
	return result
}
