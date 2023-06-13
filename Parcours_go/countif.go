package piscine

func CountIf(f func(string) bool, tab []string) int {
	cpt := 0
	for _, l := range tab {
		if f(l) {
			cpt++
		}
	}
	return cpt
}
