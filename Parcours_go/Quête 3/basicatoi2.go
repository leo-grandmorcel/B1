package piscine

func BasicAtoi2(s string) int {
	Nomber := []byte(s)
	result := 0
	grandeur := 1
	for i := len(Nomber) - 1; i >= 0; i-- {
		if int(Nomber[i]) > 47 && int(Nomber[i]) < 58 {
			result += int((Nomber[i] - 48)) * grandeur
			grandeur = grandeur * 10
		} else {
			return 0
		}
	}
	return result
}
