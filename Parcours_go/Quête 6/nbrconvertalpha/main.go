package main

import (
	"os"

	"github.com/01-edu/z01"
)

func Atoi(s string) int {
	Nomber := []byte(s)
	result := 0
	grandeur := 1
	IsNeg := false
	if len(s) == 0 {
		return 0
	}
	if int(Nomber[0]) == 45 {
		Nomber[0] = '0'
		IsNeg = true
	}
	if int(Nomber[0]) == 43 {
		Nomber[0] = '0'
	}
	for i := len(Nomber) - 1; i >= 0; i-- {
		if int(Nomber[i]) < 48 || int(Nomber[i]) > 57 {
			return 0
		} else {
			result += int((Nomber[i] - 48)) * grandeur
			grandeur = grandeur * 10
		}
	}
	if IsNeg {
		result = result * -1
	}
	return result
}

func RemoveIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

func main() {
	argument := os.Args[1:]
	str := ""
	var result []string
	var resultint []int
	maj := false
	if len(argument) == 0 {
	} else {
		for i := 0; i < len(argument); i++ {
			str = string(argument[i])
			result = append(result, str)
		}
		if result[0] == "--upper" {
			maj = true
			result = RemoveIndex(result, 0)
		}
		for i := 0; i < len(result); i++ {
			resultint = append(resultint, Atoi(result[i]))
		}
		for i := 0; i < len(resultint); i++ {
			if resultint[i] > 0 && resultint[i] < 27 {
				if maj {
					z01.PrintRune(rune(resultint[i]) + 64)
				} else {
					z01.PrintRune(rune(resultint[i]) + 96)
				}
			} else {
				z01.PrintRune(rune(32))
			}
		}
		z01.PrintRune('\n')
	}
}
