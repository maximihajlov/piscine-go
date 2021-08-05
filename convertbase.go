package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return numToBase(baseToNum(nbr, baseFrom), baseTo)
}

func numToBase(nbr int, base string) string {
	basArr := []rune(base)
	baseN := len(basArr)
	var ans []rune
	s := ""

	for nbr != 0 {
		ans = append(ans, basArr[nbr%baseN])
		nbr /= baseN
	}
	for i := range ans {
		s += string(ans[len(ans)-i-1])
	}
	return s
}

func baseToNum(s string, base string) int {
	basArr := []rune(base)
	baseN := len(basArr)
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
