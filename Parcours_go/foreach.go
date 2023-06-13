package piscine

func ForEach(f func(int), a []int) {
	for _, nombre := range a {
		f(nombre)
	}
}
