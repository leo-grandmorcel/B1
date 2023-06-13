package piscine

import "fmt"

func Game(mot string, motchache []string, n int) {
	var input string
	fmt.Println("Good Luck, you have ", n, " attempts.")
	for n > 0 {
		Affichagemot(motchache)
		fmt.Println("Press STOP to exit")
		fmt.Print("CHOOSE:  ")
		fmt.Scanln(&input)
		fmt.Print("\033[2J")
		if input == "STOP" {
			break
		}
		input = ToUpper(input)
		if len(input) >= 2 {
			if input == mot {
				fmt.Println("Congrats !")
				fmt.Println("The word was ", mot)
				n = 0
			} else {
				if n >= 2 {
					n -= 2
					Affichage((10 - (n + 1)), input, mot)
				} else {
					n = 0
					Affichage(10-(n+1), input, mot)
				}
			}
		} else if VerifLettre(input, &mot, &motchache) {
			temp := ""
			for _, letter := range motchache {
				temp += letter
			}
			if temp == mot {
				fmt.Println("Congrats !")
				fmt.Println("The word was ", mot)
				n = 0
			} else {
				fmt.Println("You find a letter")
			}
		} else {
			n -= 1
			if n == 0 {
				Affichage(10-(n+1), input, mot)
			} else {
				Affichage(10-(n+1), input, mot)
			}
		}
	}
}
