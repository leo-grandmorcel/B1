package main

import (
	"fmt"
	"os"
)

func main() {
	var name string
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		name = os.Args[1]
		file, err := os.Open(name)
		arr := make([]byte, 14)
		if err != nil {
			fmt.Println("File name missing")
		} else {
			file.Read(arr)
			fmt.Println(string(arr))
			file.Close()
		}
	}
}
