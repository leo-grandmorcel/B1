package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	str := ""
	var result []string
	for i := 1; i < len(argument); i++ {
		str = string(argument[i])
		result = append(result, str)
	}
	for i := 0; i < len(result); i++ {
		minindex := i
		for nombre := i; nombre < len(result); nombre++ {
			if result[minindex] > result[nombre] {
				minindex = nombre
			}
		}
		result[i], result[minindex] = result[minindex], result[i]
	}
	for i := 0; i < len(result); i++ {
		bonjour := []rune(result[i])
		for i := 0; i < len(bonjour); i++ {
			z01.PrintRune(bonjour[i])
		}
		z01.PrintRune('\n')
	}
}
