package piscine

func Map(f func(int) bool, a []int) []bool {
	result := []bool{}
	for _, n := range a {
		if f(n) {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
