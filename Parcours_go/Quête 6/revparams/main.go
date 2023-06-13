package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	str := ""

	for i := len(argument) - 1; i > 0; i-- {
		str = string(argument[i])
		result := []rune(str)
		for i := 0; i < len(str); i++ {
			z01.PrintRune(result[i])
		}
		z01.PrintRune('\n')
	}
}
