package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	var tmpn int = n

	if tmpn < 0 {
		z01.PrintRune('-')
		tmpn *= -1
	}
	var s [19]int8

	for i := 0; i < 19; i++ {
		if tmpn == 0 {
			s[18-i] = -1
		} else {
			s[18-i] = int8(tmpn % 10)
			tmpn /= 10
		}
	}

	for i := 0; i < 19; i++ {
		if s[i] != -1 {
			for l := '0'; l <= '9'; l++ {
				if s[i] == int8(int(l)-48) {
					z01.PrintRune(l)
				}
			}
		}
	}
	z01.PrintRune('\n')
}
