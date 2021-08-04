package Q5

func AtoiBase(s string, base string) int {
	if len(base) < 2 {
		return 0
	}
	basArr := []rune(base)
	baseN := len(basArr)
	for i := range basArr {
		for j := i + 1; j < len(basArr); j++ {
			if basArr[i] == basArr[j] || basArr[i] == '-' || basArr[i] == '+' {
				return 0
			}
		}
	}
	sArr := []rune(s)
	wNum := 0

	for i := range sArr {
		for j := 0; j < len(basArr); j++ {
			if sArr[len(sArr)-i-1] == basArr[j] {
				pow := 1
				for k := 0; k < i; k++ {
					pow *= baseN
				}
				wNum += j * pow
			}
		}
	}
	return wNum
}
