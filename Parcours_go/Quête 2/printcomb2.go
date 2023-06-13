package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb2() {
	var cpt int = 0
	for a := 48; a <= 57; a++ {
		for b := 48; b <= 57; b++ {
			for c := 48; c <= 57; c++ {
				for d := 48; d <= 57; d++ {
					if (int(rune(a)-48)*10 + int((rune(b) - 48))) < (int(rune(c)-48)*10 + int((rune(d) - 48))) {
						if cpt > 0 {
							z01.PrintRune(',')
							z01.PrintRune(rune(32))
						}
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(rune(32))
						z01.PrintRune(rune(c))
						z01.PrintRune(rune(d))
						cpt++
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
