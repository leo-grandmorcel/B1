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
		mot := piscine.Chooseword(os.Args[1])
		n := 10
		stockage := []string{}
		piscine.Game(mot, piscine.Display(mot), n, stockage)
	}
}
