package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(nbr int, base string) {
	if len(base) < 2 {
		z01.PrintRune('N')
		z01.PrintRune('V')
		return
	}
	basArr := []rune(base)
	baseN := len(basArr)
	var ans []rune
	for i := range basArr {
		for j := i + 1; j < len(basArr); j++ {
			if basArr[i] == basArr[j] || basArr[i] == '-' || basArr[i] == '+' {
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
			}
		}
	}

	if nbr < 0 {
		z01.PrintRune('-')
	}
	for nbr != 0 {
		if nbr < 0 {
			ans = append(ans, basArr[-(nbr%-baseN)])
			nbr /= -baseN
		} else {
			ans = append(ans, basArr[nbr%baseN])
			nbr /= baseN
		}
	}
	for i := len(ans) - 1; i >= 0; i-- {
		z01.PrintRune(ans[i])
	}
}
