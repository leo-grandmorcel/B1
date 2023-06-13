package piscine

func Index(s string, toFind string) int {
	sC := []byte(s)
	tC := []byte(toFind)
	cpt := 0
	var k int
	if len(s) <= 0 || len(toFind) <= 0 {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if sC[i] == tC[0] {
			k = i
			for j := 0; j < len(toFind) && j < len(s)-i; j++ {
				if !(sC[i+j] == tC[j]) {
					cpt += 1
				}
			}
			if cpt == 0 {
				return k
			} else {
				return -1
			}
		}
	}
	return -1
}
