package piscine

func BasicAtoi(s string) int {
	Nomber := []byte(s)
	result := 0
	grandeur := 1
	for i := len(Nomber) - 1; i >= 0; i-- {
		result += int((Nomber[i] - 48)) * grandeur
		grandeur = grandeur * 10
	}
	return result
}
