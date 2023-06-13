package piscine

func StrLen(s string) int {
	r := []rune(s)
	cpt := 0
	for cpt = range r {
		cpt = cpt + 1
	}
	return cpt
}
