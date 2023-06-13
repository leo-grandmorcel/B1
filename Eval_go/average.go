package piscine

func Average(tab []int) int {
	somme := 0
	cpt := 0
	result := 0
	for _, nombre := range tab {
		if nombre != -1 {
			somme += nombre
			cpt++
		}
	}
	result = somme / cpt
	return result
}
