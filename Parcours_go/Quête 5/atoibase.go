package piscine

func validBasee(base string) bool {
	var test []string
	if len(base) == 1 {
		return false
	}
	for k := 0; k < len([]rune(base)); k++ {
		if k == 0 && len(test) == 0 {
			test = append(test, string([]rune(base)[k]))
		} else {
			for x := 0; x < len(test); x++ {
				if test[x] == string([]rune(base)[k]) {
					return false
				}
			}
			test = append(test, string([]rune(base)[k]))
		}
	}
	return true
}

func AtoiBase(s string, base string) int {
	var number []int
	var temp int
	if !validBasee(base) || base[0] == '-' {
		return 0
	} else {
		for _, nb := range s {
			for i, p := range base {
				if p == nb {
					number = append(number, i)
				}
			}
		}
		temp = number[0]
		for x := 1; x < len(number); x++ {
			temp *= len(base)
			temp += number[x]
		}
	}
	return temp
}
