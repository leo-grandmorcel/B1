package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	var cpt int = 0
	for a := 48; a <= 57; a++ {
		for b := 48; b <= 57; b++ {
			for c := 48; c <= 57; c++ {
				if b > a {
					if c > b {
						if cpt > 0 {
							z01.PrintRune(rune(','))
							z01.PrintRune(rune(32))
						}
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(rune(c))
						cpt++
					}
				}
			}
		}
	}
}
