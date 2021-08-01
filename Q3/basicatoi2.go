package piscine

func BasicAtoi2(s string) int {
	sArr := []rune(s)
	n := 0
	for i := 0; i < len(sArr); i++ {
		if sArr[i] > '9' || sArr[i] < '0' {
			return 0
		}
		n *= 10
		n += int(sArr[i] - '0')
	}
	return n
}
