package piscine

func returnNbrBase(nbr int, base string) string {
	var s []int
	var result string
	temp := nbr
	for i := 0; temp > 0; i++ {
		s = append(s, temp%len(base))
		temp /= len(base)
	}
	for x := len(s) - 1; x >= 0; x-- {
		result += string(rune(base[s[x]]))
	}
	return result
}

func AtoiBase2(s string, base string) int {
	var number []int
	var temp int
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
	return temp
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	nb := AtoiBase2(nbr, baseFrom)
	result := returnNbrBase(nb, baseTo)
	return result
}
