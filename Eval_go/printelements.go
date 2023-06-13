package piscine

func PrintElements(tab []int, str []string) string {
	result := ""
	for i, nombre := range tab {
		for index, mot := range str {
			if nombre > len(str)-1 {
				if i == len(tab)-1 {
					result += "!"
				} else {
					result += "! "
				}
				break
			} else if nombre == index {
				if i == len(tab)-1 {
					result += mot
				} else {
					result += mot
					result += " "
				}
			}
		}
	}
	result += "\n"
	return result
}
