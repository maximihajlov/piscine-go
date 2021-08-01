package piscine

func BasicAtoi(s string) int {
	sArr := []rune(s)
	n := 0
	for i := 0; i < len(sArr); i++ {
		n *= 10
		n += int(sArr[i] - '0')
	}
	return n
}
