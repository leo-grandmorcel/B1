package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	var exit bool
	cpt := 0
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
		fmt.Print(result)
	} else {
		for i := 0; i < len(os.Args[3:]); i++ {
			name = os.Args[i+3]
			file, err := os.Open(name)
			if err != nil || os.Args[1] != "-c" {
				fmt.Println(err.Error())
				exit = true
				cpt++
			} else {
				longueur := Atoi(os.Args[2])
				arr := make([]byte, 100000)
				file.Read(arr)
				result := []byte{}
				for _, nombre := range arr {
					if nombre != 0 {
						result = append(result, nombre)
					}
				}
				if longueur > len(result) {
					longueur = len(result)
				}
				if cpt > 0 {
					fmt.Print("\n")
				}
				fmt.Print("==> ")
				fmt.Print(name)
				fmt.Println(" <==")
				for i := len(result) - longueur; i < len(result); i++ {
					fmt.Print(string(result[i]))
				}
				cpt++
			}
		}
	}
	if exit {
		os.Exit(1)
	}
}

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
