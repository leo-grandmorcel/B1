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

type HangManData struct {
	Word     string   `json: "Word"`
	ToFind   []string `json: "ToFind"`
	Life     int      `json:"Life"`
	Stockage []string `json: "Stockage"`
	Format   string   `json:"Format"`
}

type RankData struct {
	Username string `json : "Username"`
	Score    int    `json : "Score"`
	Parties  int    `json:"Parties"`
}

func ToUpper(s string) string {
	temp := ""
	for _, lettre := range s {
		if lettre == 'à' || lettre == 'â' || lettre == 'ä' || lettre == 'À' || lettre == 'Â' || lettre == 'Ä' {
			temp += "A"
		} else if lettre == 'é' || lettre == 'è' || lettre == 'ê' || lettre == 'ë' || lettre == 'É' || lettre == 'È' || lettre == 'Ê' || lettre == 'Ë' {
			temp += "E"
		} else if lettre == 'ï' || lettre == 'î' || lettre == 'Î' || lettre == 'Ï' {
			temp += "I"
		} else if lettre == 'ô' || lettre == 'ö' || lettre == 'Ô' || lettre == 'Ö' {
			temp += "O"
		} else if lettre == 'ù' || lettre == 'û' || lettre == 'ü' || lettre == 'Ù' || lettre == 'Û' || lettre == 'Ü' {
			temp += "U"
		} else if lettre == 'ÿ' {
			temp += "Y"
		} else if lettre == 'ç' || lettre == 'Ç' {
			temp += "C"
		} else {
			temp += string(lettre)
		}
	}
	word := []rune(temp)
	var result string = ""
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

func Display(Word string) []string {
	life := len(Word)/2 - 1
	nliste := []int{}
	str := []string{}
	Word = ToUpper(Word)
	for i := 0; life != 0; i++ {
		for j := 0; j < len(Word); j++ {
			test := true
			index := AleatoireNbr(len(Word) - 1)
			for k := 0; k < len(nliste); k++ {
				if index == nliste[k] {
					test = false
					break
				}
			}
			if test {
				nliste = append(nliste, index)
				life--
				break
			}
		}
	}
	nliste = SortIntegerTable(nliste)
	for index, lettre := range Word {
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

func Chooseword(fichier string) string {
	f, _ := os.OpenFile(fichier, os.O_RDWR, 0644)
	scanner := bufio.NewScanner(f)
	listemots := []string{}
	for scanner.Scan() {
		listemots = append(listemots, scanner.Text())
	}
	nombrealea := AleatoireNbr(len(listemots) - 1)
	return ToUpper(listemots[nombrealea])
}

func AleatoireNbr(life int) int {
	aleaint := 0
	if life >= 1 {
		rand.Seed(time.Now().UnixNano())
		aleaint = rand.Intn(life) + 1
	}
	return aleaint
}

func Affichagemot(s []string, Format string) {
	f, _ := os.OpenFile(Format, os.O_RDWR, 0644)
	scanner := bufio.NewScanner(f)
	listemots := []string{}
	for scanner.Scan() {
		listemots = append(listemots, scanner.Text())
	}
	for i := 0; i < 7; i++ {
		for j := 0; j < len(s); j++ {
			life := int((rune(s[j][0])))
			if life == 95 {
				life = 570 + i
			} else {
				life = i + int(298+(life-65)*9)
			}
			fmt.Print(listemots[life])
			fmt.Print("   ")
		}
		fmt.Println()
	}
}

func Affichage(life int, input string, Word string, Format string) {
	file, err := os.Open("affichage/hangman.txt")
	if err != nil {
		fmt.Printf("L'erreur est : %v\n", err.Error())
	}
	arr := make([]byte, 10000)
	file.Read(arr)
	file.Close()
	temp := ((7 * 11) * life) + (life * 2)
	if len(input) > 1 {
		if 10-(life+1) != 0 {
			fmt.Println("Wrong word,", 10-(life+1), "attempts remaining")
		}
	} else {
		if 10-(life+1) != 0 {
			fmt.Println(input, "is not present in the word, ", 10-(life+1), " attempts remaining")
		}
	}
	for i := temp; i <= temp+78; i++ {
		fmt.Print(string(arr[i]))
	}
	if 10-(life+1) == 0 {
		str := []string{}
		for i := 0; i < len(Word); i++ {
			str = append(str, string(Word[i]))
		}
		Affichagemot([]string{"J", "O", "S", "E", " ", "I", "S", " ", "H", "A", "N", "G", "E", "D"}, Format)
		Affichagemot(str, Format)
	}
}

func SameLetter(input string, stockage *[]string) bool {
	for _, lettre := range *stockage {
		if input == lettre {
			return false
		}
	}
	*stockage = append(*stockage, input)
	return true
}

func VerifLettre(let string, Word *string, ToFind *[]string) bool {
	LtrFind := false
	for index, lettre := range *Word {
		if let == string(lettre) {
			if (*ToFind)[index] != let {
				(*ToFind)[index] = let
				LtrFind = true
			}
		}
	}
	return LtrFind
}

func DeleteSave(savename string) {
	if _, err := os.Stat(savename); err == nil {
		os.Remove(savename)
	}
}

func Code(data HangManData, name string) {
	b, _ := json.Marshal(data)
	f, _ := os.Create(name)
	f.Write(b)
	f.Close()
}

func Decode(savename string) HangManData {
	f, err := os.OpenFile(savename, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	b, _ := os.ReadFile(savename)
	dataSave := HangManData{}
	_ = json.Unmarshal(b, &dataSave)
	f.Close()
	return dataSave
}

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func CallClear() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func OpenRank(rankname string, score int) {
	rankname = ToUpper(rankname)
	f, err := os.OpenFile("rank.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	tab := []RankData{}
	for scanner.Scan() {
		dRank := RankData{}
		_ = json.Unmarshal(scanner.Bytes(), &dRank)
		tab = append(tab, dRank)
	}
	f.Close()
	test := true
	for index, player := range tab {
		if player.Username == rankname {
			tab[index].Score += score
			tab[index].Parties++
			test = false
			break
		}
	}
	if test {
		newuser := RankData{rankname, score, 1}
		tab = append(tab, newuser)
	}
	for index := range tab {
		minindex := index
		for i := index; i < len(tab); i++ {
			if tab[minindex].Score < tab[i].Score {
				minindex = i
			}
		}
		tab[index], tab[minindex] = tab[minindex], tab[index]
	}
	f, _ = os.OpenFile("rank.txt", os.O_RDWR|os.O_TRUNC, 0777)
	for index := range tab {
		b, _ := json.Marshal(tab[index])
		f.Write(b)
		f.WriteString("\n")
	}
	f.Close()
	AfficheRank()
}

func AfficheRank() {
	fmt.Println("USERNAME             SCORE     GAMES")
	f, err := os.OpenFile("rank.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	tab := []RankData{}
	for scanner.Scan() {
		dRank := RankData{}
		_ = json.Unmarshal(scanner.Bytes(), &dRank)
		tab = append(tab, dRank)
	}
	f.Close()
	for i := range tab {
		long := 21 - len(tab[i].Username)
		fmt.Print(tab[i].Username)
		for j := 0; j < long; j++ {
			fmt.Print(" ")
		}
		fmt.Print(tab[i].Score)
		long = 10 - LenNbr(tab[i].Score)
		for j := 0; j < long; j++ {
			fmt.Print(" ")
		}
		fmt.Println(tab[i].Parties)
	}
}

func LenNbr(n int) int {
	if n == 0 {
		return 1
	}
	var t []int
	for n != 0 {
		reste := n % 10
		t = append(t, reste)
		n = (n / 10)
	}
	return len(t)
}
