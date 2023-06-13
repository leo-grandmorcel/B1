package piscine

func ConcatParams(args []string) string {
	var result string
	for i := 0; i < len(args); i++ {
		result += args[i]
		if i < len(args)-1 {
			result += "\n"
		}
	}
	return result
}
