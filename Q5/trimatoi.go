package Q5

func TrimAtoi(s string) int {
	sArr := []rune(s)
	ans := 0
	minusWas := false
	for i := range sArr {
		if '0' <= sArr[i] && sArr[i] <= '9' {
			ans *= 10
			ans += int(sArr[i] - '0')
		} else if sArr[i] == '-' && !minusWas && ans == 0 {
			minusWas = true
		}
	}
	if minusWas {
		ans *= -1
	}
	return ans
}
