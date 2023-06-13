package piscine

func AppendRange(min, max int) []int {
	result := []int(nil)
	for i := min; i < max; i++ {
		result = append(result, i)
	}
	return result
}
