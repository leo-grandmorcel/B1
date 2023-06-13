package piscine

func MakeRange(min, max int) []int {
	size := max - min
	if size <= 0 {
		return []int(nil)
	} else {
		cpt := 0
		result := make([]int, size)
		for i := min; i < max; i++ {
			result[cpt] = i
			cpt++
		}
		return result
	}
}
