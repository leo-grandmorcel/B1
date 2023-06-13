package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	result := true
	if len(a) == 1 {
	} else {
		if a[0] > a[1] {
			for i := 0; i < len(a); i++ {
				if i == len(a)-1 {
					continue
				} else if f(a[i], a[i+1]) >= 0 {
					result = true
				} else {
					return false
				}
			}
		} else {
			for i := len(a) - 1; i >= 0; i-- {
				if i == 0 {
					continue
				} else if f(a[i], a[i-1]) >= 0 {
					result = true
				} else {
					return false
				}
			}
		}
	}
	return result
}
