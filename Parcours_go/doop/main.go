package main

import (
	"os"
)

func main() {
	if len(os.Args[1:]) == 3 {
		valeur := 0
		value1 := os.Args[1]
		value2 := os.Args[3]
		operator := os.Args[2]
		if value1 == "9223372036854775809" || value2 == "9223372036854775809" {
		} else {
			value1 := Atoi(os.Args[1])
			value2 := Atoi(os.Args[3])
			if value1 >= 9223372036854775807 || value2 >= 9223372036854775807 {
			} else {
				if operator == "+" {
					valeur = value1 + value2
					InttoString(valeur)
				} else if operator == "-" {
					valeur = value1 - value2
					InttoString(valeur)
				} else if operator == "/" {
					if value2 == 0 {
						os.Stderr.WriteString("No division by 0")
						os.Stderr.WriteString("\n")
					} else {
						valeur = value1 / value2
						if valeur == 0 {
							result := []byte{48}
							os.Stdout.Write(result)
						}
						InttoString(valeur)
					}
				} else if operator == "*" {
					valeur = value1 * value2
					InttoString(valeur)
				} else if operator == "%" {
					if value2 == 0 {
						os.Stderr.WriteString("No modulo by 0")
						os.Stderr.WriteString("\n")
					} else {
						valeur = value1 % value2
						InttoString(valeur)
					}
				}
			}
		}
	}
}

func Atoi(s string) int {
	Nomber := []byte(s)
	result := 0
	grandeur := 1
	IsNeg := false
	if s == "9223372036854775808" {
		return 9223372036854775807
	}
	if len(s) == 0 {
		return 0
	}
	if int(Nomber[0]) == 45 {
		Nomber[0] = '0'
		IsNeg = true
	}
	if int(Nomber[0]) == 43 {
		Nomber[0] = '0'
	}
	for i := len(Nomber) - 1; i >= 0; i-- {
		if int(Nomber[i]) < 48 || int(Nomber[i]) > 57 {
			return 9223372036854775807
		} else {
			result += int((Nomber[i] - 48)) * grandeur
			grandeur = grandeur * 10
		}
	}
	if IsNeg {
		result = result * -1
	}
	return result
}

func InttoString(nbr int) {
	var s []int
	if nbr < 0 {
		nbr *= -1
		os.Stderr.WriteString("-")
	}
	temp := nbr
	for i := 0; temp > 0; i++ {
		s = append(s, temp%10)
		temp /= 10
	}
	var result []byte
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, byte(s[i]+48))
	}
	os.Stdout.Write(result)
	os.Stderr.WriteString("\n")
}
