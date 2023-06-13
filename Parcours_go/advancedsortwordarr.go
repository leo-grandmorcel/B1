package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for index := 0; index+1 < len(a)-1; index++ {
		minindex := index
		for nombre := index; nombre < len(a); nombre++ {
			if f((a[minindex]), a[nombre]) == -1 {
				minindex = nombre
			}
		}
		a[index], a[minindex] = a[minindex], a[index]
	}
}
