package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	if len(os.Args[1:]) <= 0 {
		fmt.Println("No given file")
		fmt.Println("If you want to start a new game give a file.txt in argument")
		fmt.Println("If you want to start at a save give the flag --startWith and file.txt")
	} else {
		if os.Args[1] == "--startWith" {
			if len(os.Args[1:]) == 1 {
				fmt.Println("No save.txt given")
			} else {
				mot, motvide, n, stockage := piscine.Decode(os.Args[2])
				piscine.Game(mot, motvide, n, stockage, os.Args[2])
			}
		} else {
			mot := piscine.Chooseword(os.Args[1])
			n := 10
			stockage := []string{}
			piscine.Game(mot, piscine.Display(mot), n, stockage, "")
		}
	}
}
