package piscine

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func ToUpper(s string) string {
	temp := ""
	for _, lettre := range s {
		if lettre == 'à' || lettre == 'â' || lettre == 'ä' {
			temp += "A"
		} else if lettre == 'é' || lettre == 'è' || lettre == 'ê' || lettre == 'ë' {
			temp += "E"
		} else if lettre == 'ï' || lettre == 'î' {
			temp += "I"
		} else if lettre == 'ô' || lettre == 'ö' {
			temp += "O"
		} else if lettre == 'ù' || lettre == 'û' || lettre == 'ü' {
			temp += "U"
		} else if lettre == 'ÿ' {
			temp += "Y"
		} else if lettre == 'ç' {
			temp += "C"
		} else {
			temp += string(lettre)
		}
	}
	word, result := []rune(temp), ""
	for index := range temp {
		if word[index] >= 97 && word[index] <= 122 {
			word[index] = word[index] - 32
			result += string(word[index])
		} else if word[index] >= 65 && word[index] <= 90 {
			result += string(word[index])
		}
	}
	return result
}

func SortIntegerTable(table []int) []int {
	for index := 0; index < len(table)-1; index++ {
		minindex := index
		for nombre := index; nombre < len(table); nombre++ {
			if table[minindex] > table[nombre] {
				minindex = nombre
			}
		}
		table[index], table[minindex] = table[minindex], table[index]
	}
	return table
}

func Display(s string) []string {
	n, nliste, str := len(s)/2-1, []int{}, []string{}
	s = ToUpper(s)
	for i := 0; n != 0; i++ {
		for j := 0; j < len(s); j++ {
			index, test := AleatoireNbr(len(s)-1), true
			for k := 0; k < len(nliste); k++ {
				if index == nliste[k] {
					test = false
					break
				}
			}
			if test {
				nliste = append(nliste, index)
				n--
				break
			}
		}
	}
	nliste = SortIntegerTable(nliste)
	for index, lettre := range s {
		test := true
		for _, nombre := range nliste {
			if index == nombre {
				str = append(str, string(lettre))
				test = false
				break
			}
		}
		if test {
			str = append(str, "_")
		}
	}
	return str
}

func Chooseword(s string) string {
	f, _ := os.OpenFile(s, os.O_RDWR, 0644)
	scanner, listemots := bufio.NewScanner(f), []string{}
	for scanner.Scan() {
		listemots = append(listemots, scanner.Text())
	}
	nombrealea := AleatoireNbr(len(listemots) - 1)
	return ToUpper(listemots[nombrealea])
}

func AleatoireNbr(n int) int {
	aleaint := 0
	if n >= 1 {
		rand.Seed(time.Now().UnixNano())
		aleaint = rand.Intn(n) + 1
	}
	return aleaint
}

func Affichagemot(s []string) {
	for i := 0; i < len(s); i++ {
		if i == 0 {
			fmt.Print(s[i])
		} else {
			fmt.Print(" ", s[i])
		}
	}
	fmt.Println()
	fmt.Println()
}

func Affichage(n int, input string, mot string) {
	file, err := os.Open("hangman.txt")
	if err != nil {
		fmt.Printf("L'erreur est : %v\n", err.Error())
	}
	arr, temp := make([]byte, 10000), ((7*11)*n)+(n*2)
	file.Read(arr)
	file.Close()
	if len(input) > 1 {
		if 10-(n+1) != 0 {
			fmt.Println("Wrong word,", 10-(n+1), "attempts remaining")
		}
	} else {
		if 10-(n+1) != 0 {
			fmt.Println(input, "is not present in the word, ", 10-(n+1), " attempts remaining")
		}
	}
	for i := temp; i <= temp+78; i++ {
		fmt.Print(string(arr[i]))
	}
	if 10-(n+1) == 0 {
		fmt.Println("José is hanged")
		fmt.Println("The word was", mot)
	}
}

func VerifLettre(let string, mot *string, motchache *[]string) bool {
	bon := false
	for index, lettre := range *mot {
		if let == string(lettre) {
			if (*motchache)[index] != let {
				(*motchache)[index] = let
				bon = true
			}
		}
	}
	return bon
}
