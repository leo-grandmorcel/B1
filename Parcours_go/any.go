package piscine

func Any(f func(string) bool, a []string) bool {
	for _, n := range a {
		if f(n) {
			return true
		}
	}
	return false
}
