package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase2(nbr int, base string) {
	str := []rune(base)
	IntNeg := false
	stock := nbr
	table := []int{}
	for i := 0; i < len(base); i++ {
		for j := i + 1; j < len(base); j++ {
			if (str[i] == str[j]) || (str[i] == 43) || (str[i] == 45) || (-999999999999999 > nbr) || (nbr > 999999999999) {
				return
			}
		}
	}
	if stock < 0 {
		IntNeg = true
		stock = -stock
	}
	for i := len(base); stock >= 1; {
		stock, table = stock/i, append(table, stock%i)
	}
	if IntNeg {
		z01.PrintRune(45)
	}
	for i := len(table) - 1; i >= 0; i-- {
		z01.PrintRune(rune(str[table[i]]))
	}
}
