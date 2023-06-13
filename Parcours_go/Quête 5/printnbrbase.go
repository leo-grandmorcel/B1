package piscine

import (
	"fmt"

	"github.com/01-edu/z01"
)

func validBase(base string) bool {
	var test []string
	if len(base) == 1 {
		return false
	}
	fmt.Println("longeur table rune base", []rune(base))
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

func PrintNbrBase(nbr int, base string) {
	var s []int
	if nbr == -9223372036854775808 {
		number := "-9223372036854775808"
		for _, nb := range number {
			z01.PrintRune(nb)
		}
	} else {
		if !validBase(base) || base[0] == '-' {
			z01.PrintRune('N')
			z01.PrintRune('V')
		} else {
			if nbr < 0 {
				nbr *= -1
				z01.PrintRune('-')
			}
			temp := nbr
			for i := 0; temp > 0; i++ {
				s = append(s, temp%len(base))
				temp /= len(base)
			}
		}
		for x := len(s) - 1; x >= 0; x-- {
			z01.PrintRune(rune(base[s[x]]))
		}
	}
}
