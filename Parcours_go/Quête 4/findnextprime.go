package piscine

func FindNextPrime(nb int) int {
	if nb <= 0 {
		return 2
	} else if IsPrime(nb) {
		return nb
	} else {
		return FindNextPrime(nb + 1)
	}
}
