package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	sArr := []rune(s)
	minus := 1
	if sArr[0] == '-' {
		minus = -1
		sArr[0] = '0'
	} else if sArr[0] == '+' {
		sArr[0] = '0'
	}
	n := 0
	for i := 0; i < len(sArr); i++ {
		if sArr[i] < '0' || '9' < sArr[i] {
			return 0
		}
		n *= 10
		n += int(sArr[i] - '0')
	}
	return minus * n
}
