package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div, *mod = a/b, a%b
}
