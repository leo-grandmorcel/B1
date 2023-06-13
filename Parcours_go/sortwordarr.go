package piscine

func SortWordArr(a []string) {
	for index := 0; index < len(a)-1; index++ {
		minindex := index
		for nombre := index; nombre < len(a); nombre++ {
			if a[minindex] > a[nombre] {
				minindex = nombre
			}
		}
		a[index], a[minindex] = a[minindex], a[index]
	}
}
