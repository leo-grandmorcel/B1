package main

import (
	"fmt"
	"os"
	"piscine"
)

func main() {
	if len(os.Args[1:]) <= 0 {
		fmt.Println("No given file")
	} else {
		mot, n := piscine.Chooseword(os.Args[1]), 10
		fmt.Println(mot)
		piscine.Game(mot, piscine.Display(mot), n)
	}
}
