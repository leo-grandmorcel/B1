package piscine

import (
	"fmt"
)

func Game(mot string, motchache []string, n int, stockage []string, loadsave string) {
	var input string
	CallClear()
	fmt.Println("Good Luck, you have ", n, " attempts.")
	for n > 0 {
		Affichagemot(motchache)
		fmt.Println("Press STOP to exit")
		fmt.Print("CHOOSE:  ")
		fmt.Scanln(&input)
		CallClear()
		if input == "STOP" {
			var namesave string
			fmt.Print("Name of your save :  ")
			fmt.Scanln(&namesave)
			Code(&mot, &motchache, &n, &stockage, (namesave + ".txt"))
			fmt.Printf("Game save in %s.txt", namesave)
			break
		}
		input = ToUpper(input)
		if len(input) >= 2 {
			if input == mot {
				fmt.Println("Congrats !")
				fmt.Println("The word was ", mot)
				DeleteSave(loadsave)
				n = 0
			} else {
				if n >= 2 {
					n -= 2
					Affichage((10 - (n + 1)), input, mot)
				} else {
					n = 0
					Affichage(10-(n+1), input, mot)
					DeleteSave(loadsave)
				}
			}
		} else if Memelettre(input, &stockage) {
			if VerifLettre(input, &mot, &motchache) {
				temp := ""
				for _, letter := range motchache {
					temp += letter
				}
				if temp == mot {
					fmt.Println("Congrats !")
					fmt.Println("The word was ", mot)
					DeleteSave(loadsave)
					n = 0
				} else {
					fmt.Println("You find a letter")
				}
			} else {
				n -= 1
				if n == 0 {
					Affichage(10-(n+1), input, mot)
					DeleteSave(loadsave)
				} else {
					Affichage(10-(n+1), input, mot)
				}
			}
		} else {
			fmt.Println("You have already found this letter")
		}
	}
}
