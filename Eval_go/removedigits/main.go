package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	cpt := 1
	test := true
	format := "Argument "
	if len(os.Args[1:]) >= 1 {
		argument := os.Args[1:]
		for i := 0; i < len(argument); i++ {
			for j := 0; j < len(argument[i]); j++ {
				if argument[i][j] >= 48 && argument[i][j] <= 57 {
					test = false
					break
				}
			}
			if test {
				for k := 0; k < len(format); k++ {
					z01.PrintRune(rune(format[k]))
				}
				z01.PrintRune(rune(i + 49))
				z01.PrintRune(32)
				z01.PrintRune(':')
				z01.PrintRune(32)
				for j := 0; j < len(argument[i]); j++ {
					z01.PrintRune(rune(argument[i][j]))

				}
				z01.PrintRune('\n')
				cpt++
			}
			test = true
		}
	} else {
		z01.PrintRune('\n')
	}
}
