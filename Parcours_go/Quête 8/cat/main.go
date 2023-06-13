package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	var name string
	if len(os.Args) == 1 {
		bytes, _ := os.Open(os.Stdin.Name())
		arr := make([]byte, 100000)
		bytes.Read(arr)
		result := []byte{}
		for _, letter := range arr {
			if letter != 0 {
				result = append(result, letter)
			}
		}
		for _, letter := range result {
			z01.PrintRune(rune(letter))
		}
	} else {
		for i := 0; i < len(os.Args[1:]); i++ {
			name = os.Args[i+1]
			file, err := os.Open(name)
			if err != nil {
				str := "ERROR: " + (err.Error())

				for _, letter := range str {
					z01.PrintRune(letter)
				}
				z01.PrintRune('\n')
				os.Exit(1)
			} else {
				arr := make([]byte, 100000)
				file.Read(arr)
				result := []byte{}
				for _, nombre := range arr {
					if nombre != 0 {
						result = append(result, nombre)
					}
				}

				for _, nombre := range result {
					z01.PrintRune(rune(nombre))
				}
			}
		}
	}
}
