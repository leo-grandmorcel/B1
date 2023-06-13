package piscine

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

var clear map[string]func()

type save struct {
	Mot       *string
	Motrempli *[]string
	N         *int
	Stockage  *[]string
}
type loadsave struct {
	Mot       string   `json: "Mot"`
	Motrempli []string `json: "Motrempli"`
	N         int      `json:"N"`
	Stockage  []string `json: "Stockage"`
}

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
			test := true
			index := AleatoireNbr(len(s) - 1)
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

func DeleteSave(namesave string) {
	if _, err := os.Stat(namesave); err == nil {
		os.Remove(namesave)
	}
}

func Code(mot *string, motrempli *[]string, n *int, stockage *[]string, namesave string) {
	m := save{mot, motrempli, n, stockage}
	b, _ := json.Marshal(m)
	f, _ := os.Create(namesave)
	f.Write(b)
	f.Close()
}

func Decode(nom string) (string, []string, int, []string) {
	f, err := os.OpenFile(nom, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := os.ReadFile(nom)
	m := loadsave{}
	_ = json.Unmarshal(b, &m)
	f.Close()
	return m.Mot, m.Motrempli, m.N, m.Stockage
}

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func Chooseword(s string) string {
	f, _ := os.OpenFile(s, os.O_RDWR, 0644)
	scanner := bufio.NewScanner(f)
	listemots := []string{}
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
	arr := make([]byte, 10000)
	file.Read(arr)
	file.Close()
	temp := ((7 * 11) * n) + (n * 2)
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

func Memelettre(s string, stockage *[]string) bool {
	for _, lettre := range *stockage {
		if s == lettre {
			return false
		}
	}
	*stockage = append(*stockage, s)
	return true
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
