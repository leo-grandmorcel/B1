package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Isvoyelle(l rune) bool {
	if l == 'A' || l == 'a' || l == 'E' || l == 'e' || l == 'I' || l == 'i' || l == 'O' || l == 'o' || l == 'U' || l == 'u' {
		return true
	}
	return false
}

func main() {
	argument := os.Args[1:]
	var strTemp string
	var tableV []byte
	for i := 0; i < len(argument); i++ {
		strTemp += argument[i]
		strTemp += " "
	}
	for _, i := range strTemp {
		if Isvoyelle(i) {
			tableV = append(tableV, byte(i))
		}
	}
	for _, i := range strTemp {
		if Isvoyelle(i) {
			z01.PrintRune(rune(tableV[len(tableV)-1]))
			tableV = tableV[:len(tableV)-1]
		} else {
			z01.PrintRune(i)
		}
	}
	z01.PrintRune('\n')
}
