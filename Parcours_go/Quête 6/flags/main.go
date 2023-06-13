package main

import (
	"fmt"
	"os"
)

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

func main() {
	argument := os.Args[1:]
	IsInsert := false
	IsOrder := false
	IsHelp := false
	var strInsert string
	var strFinal string
	var strTemp string
	strHelp := "--insert\n" + "  -i\n" + "\t This flag inserts the string into the string passed as argument.\n" + "--order\n" + "  -o\n" + "\t This flag will behave like a boolean, if it is called it will order the argument."
	if len(argument) == 0 {
		fmt.Println(strHelp)
	} else {
		for i := 0; i < len(argument); i++ {
			ElemArg := argument[i]
			if len(ElemArg) <= 0 {
			} else if ElemArg[0] == '-' {
				if len(ElemArg) >= 3 {
					if ElemArg[1] == 'i' || ElemArg[2] == 'i' {
						IsInsert = true
						for k := 0; k < len(ElemArg); k++ {
							if rune(ElemArg[k]) == 61 {
								strInsert = ElemArg[k+1:]
								break
							}
						}
					} else if ElemArg[1] == 'o' || ElemArg[2] == 'o' {
						IsOrder = true
					} else if ElemArg[1] == 'h' || ElemArg[2] == 'h' {
						IsHelp = true
					}
				} else {
					if ElemArg[1] == 'i' {
						IsInsert = true
						for k := 0; k < len(ElemArg); k++ {
							if rune(ElemArg[k]) == 61 {
								strInsert = ElemArg[k+1:]
							}
						}
					} else if ElemArg[1] == 'o' {
						IsOrder = true
					} else if ElemArg[1] == 'h' {
						IsHelp = true
					}
				}
			} else {
				strTemp += ElemArg
			}
		}
		if IsHelp {
			IsInsert, IsHelp, IsOrder = false, false, false
			fmt.Println(strHelp)
		} else {
			if IsInsert {
				strTemp += strInsert
			}
			if IsOrder {
				var table []int
				for i := 0; i < len(strTemp); i++ {
					table = append(table, int(strTemp[i]))
				}
				table = SortIntegerTable(table)
				for i := 0; i < len(table); i++ {
					strFinal += string(table[i])
				}
			} else {
				strFinal = strTemp
			}
			if !IsHelp {
				fmt.Println(strFinal)
			}
		}
	}
}
