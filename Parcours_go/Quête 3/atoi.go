package piscine

func Atoi(s string) int {
	Nomber := []byte(s)
	result := 0
	grandeur := 1
	IsNeg := false
	if len(s) == 0 {
		return 0
	}
	if int(Nomber[0]) == 45 {
		Nomber[0] = '0'
		IsNeg = true
	}
	if int(Nomber[0]) == 43 {
		Nomber[0] = '0'
	}
	for i := len(Nomber) - 1; i >= 0; i-- {
		if int(Nomber[i]) < 48 || int(Nomber[i]) > 57 {
			return 0
		} else {
			result += int((Nomber[i] - 48)) * grandeur
			grandeur = grandeur * 10
		}
	}
	if IsNeg {
		result = result * -1
	}
	return result
}
