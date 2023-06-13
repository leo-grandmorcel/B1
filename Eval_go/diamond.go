package piscine

import "github.com/01-edu/z01"

func Diamond(index int) {
	x := (index * 2) - 1
	y := (index * 2) - 1
	if index > 1 {
		for ligne := 1; ligne <= y; ligne++ {
			for col := 1; col <= x; col++ {
				if ligne == 1 || ligne == x {
					if col == index {
						z01.PrintRune(42)
					} else {
						z01.PrintRune(32)
					}
				} else if ligne <= 5 {
					if ligne == (index-col+1) || ligne == (col-index+1) {
						z01.PrintRune(42)
					} else {
						z01.PrintRune(32)
					}
				} else if ligne > 5 {
					templigne := ligne - index
					if templigne == col-1 || templigne == y-col {
						z01.PrintRune(42)
					} else {
						z01.PrintRune(32)
					}
				}
			}
			z01.PrintRune('\n')
		}
	} else if index > 0 {
		z01.PrintRune(42)
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('\n')
	}

}
