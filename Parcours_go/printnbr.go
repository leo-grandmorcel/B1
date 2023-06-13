package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n == -9223372036854775808 {
		number := "-9223372036854775808"
		for _, nb := range number {
			z01.PrintRune(nb)
		}
	} else {
		var s []int
		temp := n
		if n < 0 {
			temp *= -1
			z01.PrintRune('-')
		}
		if temp == 0 {
			z01.PrintRune(48)
		} else {
			for i := 0; temp > 0; i++ {
				s = append(s, temp%10)
				temp /= 10
			}
			for i := len(s) - 1; i >= 0; i-- {
				z01.PrintRune(rune(s[i]) + 48)
			}
		}
	}
}
