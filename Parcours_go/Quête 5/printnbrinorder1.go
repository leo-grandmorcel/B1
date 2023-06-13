package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	cpt := 0
	tempint := n
	for i := 10; tempint > 0; {
		tempint = tempint / i
		cpt++
	}
	m := 0
	result := []byte("")
	for i := 0; i < cpt; i++ {
		m = n / 10
		l1 := n % 10
		result = append(result, byte(l1))
		n = m
	}
	if len(result) == 0 {
		result = append(result, byte(0))
	}
	for i := 0; i < cpt; i++ {
		for j := 1; j < cpt; j++ {
			if result[j] < result[j-1] {
				result[j], result[j-1] = result[j-1], result[j]
			}
		}
	}
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]) + 48)
	}
}
