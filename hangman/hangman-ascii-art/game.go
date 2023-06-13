package piscine

import "fmt"

func Game(x HangManData, loadsave string) {
	CallClear()
	var input string
	fmt.Println("Different username :")
	fmt.Println()
	AfficheRank()
	fmt.Println()
	fmt.Print("Create or give your username :  ")
	username := ""
	for {
		fmt.Scanln(&username)
		if len(username) <= 20 {
			break
		} else {
			CallClear()
			fmt.Println("Your username is too long only 20 characters")
			fmt.Print("Create or give your username :  ")
		}
	}
	CallClear()
	fmt.Println("Good Luck, you have ", x.Life, " attempts.")
	for x.Life > 0 {
		Affichagemot(x.ToFind, x.Format)
		fmt.Println("Press STOP to exit")
		fmt.Print("CHOOSE:  ")
		fmt.Scanln(&input)
		CallClear()
		if input == "STOP" {
			var namesave string
			fmt.Print("Name of your save :  ")
			fmt.Scanln(&namesave)
			if len(namesave) > 0 {
				namesave = namesave + ".txt"
				Code(x, "saves/"+namesave)
				fmt.Printf("Game save in %s", namesave)
			} else {
				fmt.Println("Game gived up ")
			}
			break
		}
		input = ToUpper(input)
		if len(input) >= 2 {
			if input == x.Word {
				str := []string{}
				for i := 0; i < len(x.Word); i++ {
					str = append(str, string(x.Word[i]))
				}
				fmt.Println("Congrats !")
				Affichagemot(str, x.Format)
				DeleteSave(loadsave)
				OpenRank(username, x.Life*10)
				x.Life = 0
			} else {
				if x.Life >= 2 {
					x.Life -= 2
					Affichage((10 - (x.Life + 1)), input, x.Word, x.Format)
				} else {
					x.Life = 0
					Affichage(10-(x.Life+1), input, x.Word, x.Format)
					DeleteSave(loadsave)
					OpenRank(username, x.Life*10)
				}
			}
		} else if SameLetter(input, &x.Stockage) {
			if VerifLettre(input, &x.Word, &x.ToFind) {
				temp := ""
				for _, letter := range x.ToFind {
					temp += letter
				}
				if temp == x.Word {
					str := []string{}
					for i := 0; i < len(x.Word); i++ {
						str = append(str, string(x.Word[i]))
					}
					fmt.Println("Congrats !")
					Affichagemot(str, x.Format)
					DeleteSave(loadsave)
					OpenRank(username, x.Life*10)
					x.Life = 0
				} else {
					fmt.Println("You find a letter")
				}
			} else {
				x.Life -= 1
				if x.Life == 0 {
					Affichage(10-(x.Life+1), input, x.Word, x.Format)
					DeleteSave(loadsave)
					OpenRank(username, x.Life*10)
				} else {
					Affichage(10-(x.Life+1), input, x.Word, x.Format)
				}
			}
		} else {
			fmt.Println("You have already found this letter")
		}
	}
}
