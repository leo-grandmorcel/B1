package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argu := os.Args
	last := len(argu) - 1
	var s string
	s = s + string(argu[last])
	arg := []rune(s)
	var chars []rune
	for i := len(arg) - 1; i >= 0; i-- {
		if arg[i] != 92 && arg[i] != 47 {
			chars = append(chars, arg[i])
		} else {
			break
		}
	}
	for i := len(chars) - 1; i >= 0; i-- {
		z01.PrintRune(chars[i])
	}
	z01.PrintRune('\n')
}
