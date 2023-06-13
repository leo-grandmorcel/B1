package piscine

import "github.com/01-edu/z01"

func SortIntegerTable(table []int) []int {
	for index := 0; index < len(table)-1; index++ {
		minindex := index
		for nombre := index; nombre < len(table); nombre++ {
			if table[minindex] > table[nombre] {
				minindex = nombre
			}
		}
		table[index], table[minindex] = table[minindex], table[index]
	}
	return table
}

func PrintNbrInOrder2(n int) {
	cpt := 0
	tempint := n
	for i := 10; tempint > 0; cpt++ {
		tempint = tempint / i
	}
	result := []int{}
	for i := 0; i < cpt; i++ {
		m := n / 10
		l1 := n % 10
		result = append(result, int(l1))
		n = m
	}
	if len(result) == 0 {
		result = append(result, int(0))
	}
	result = SortIntegerTable(result)
	for i := 0; i < len(result); i++ {
		z01.PrintRune(rune(result[i]) + 48)
	}
}
