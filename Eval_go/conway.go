package piscine

func Conway(index int) int {
	if index > 1 {
		x := Conway(index - 1)
		if x == 0 {
			x = 1
		}
		str := PrintNbr(x)
		result := ""
		resultint := 1
		var valcpt byte
		for i := 0; i < len(str); i++ {
			if valcpt != str[i] {
				valcpt = str[i]
			} else {
				continue
			}
			cpt := 0
			for j := i; j < len(str); j++ {
				if valcpt == str[j] {
					cpt++
				} else {
					break
				}
			}
			result += PrintNbr(cpt)
			result += string(valcpt)
		}
		resultint = Atoi(result)
		return resultint
	} else {
		return 0
	}
}

func PrintNbr(n int) string {
	result := ""
	var s []int
	temp := n
	if temp == 0 {
		result += "0"
	} else {
		for i := 0; temp > 0; i++ {
			s = append(s, temp%10)
			temp /= 10
		}
		for i := len(s) - 1; i >= 0; i-- {
			result += string(rune(s[i] + 48))
		}
	}
	return result
}

func Atoi(s string) int {
	Nomber := []byte(s)
	result := 0
	grandeur := 1
	IsNeg := false
	if len(s) == 0 {
		return 0
	}
	if int(Nomber[0]) == 45 {
		Nomber[0] = '0'
		IsNeg = true
	}
	if int(Nomber[0]) == 43 {
		Nomber[0] = '0'
	}
	for i := len(Nomber) - 1; i >= 0; i-- {
		if int(Nomber[i]) < 48 || int(Nomber[i]) > 57 {
			return 0
		} else {
			result += int((Nomber[i] - 48)) * grandeur
			grandeur = grandeur * 10
		}
	}
	if IsNeg {
		result = result * -1
	}
	return result
}
