package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = 223372036854775808
		} else {
			n *= -1
		}
	}
	var s [20]int8

	for i := 0; i < 20; i++ {
		if n == 0 {
			s[19-i] = -1
		} else {
			s[19-i] = int8(n % 10)
			n /= 10
		}
	}

	for i := 0; i < 20; i++ {
		if s[i] != -1 {
			for l := '0'; l <= '9'; l++ {
				if s[i] == int8(int(l)-48) {
					z01.PrintRune(l)
				}
			}
		}
	}
}
