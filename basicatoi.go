package piscine

func BasicAtoi(s string) int {
	sArr := []rune(s)
	n := 0
	for i := 0; i < len(sArr); i++ {
		pow := 1
		for j := 1; j < len(sArr)-i; j++ {
			pow *= 10
		}

		n += int(sArr[i]-48) * pow
	}
	return n
}
