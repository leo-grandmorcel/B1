package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"piscine"
)

func main() {
	if len(os.Args[1:]) <= 0 {
		fmt.Println("No given file")
		fmt.Println("If you want to start a new game give a file.txt in argument")
		fmt.Println("If you want to start at a save give the flag --startWith and file.txt")
	} else if len(os.Args[1:]) == 3 && os.Args[2] == "--letterFile" {
		data := piscine.HangManData{}
		data.Word = piscine.Chooseword("wordslist/" + os.Args[1])
		data.ToFind = piscine.Display(data.Word)
		data.Life = 10
		data.Format = "affichage/" + os.Args[3]
		piscine.Game(data, "")
	} else if os.Args[1] == "--startWith" {
		if len(os.Args[1:]) == 1 {
			files, err := ioutil.ReadDir("saves")
			if err != nil {
				log.Fatal(err)
			}
			if len(files) <= 0 {
				fmt.Println("No save found")
			} else {
				fmt.Println("Different save in memory :")
				fmt.Println()
				for _, file := range files {
					tmp := string(file.Name())
					nb := len(tmp) - 4
					tmp = tmp[:nb]
					fmt.Println("- ", tmp)
				}
				sauv := ""
				fmt.Println()
				fmt.Print("Which save you want ?  ")
				fmt.Scanln(&sauv)
				sauv += ".txt"
				args := "saves/" + sauv
				data := piscine.Decode(args)
				piscine.Game(data, args)
			}
		}
	} else {
		fmt.Println("Not enough information")
	}
}
