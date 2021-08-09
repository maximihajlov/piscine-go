package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)
	printStr("x = ")
	PrintNbr(points.x)
	printStr(", y = ")
	PrintNbr(points.y)
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
		if n == -9223372036854775808 {
			z01.PrintRune('9')
			n = 223372036854775808
		} else {
			n *= -1
		}
	} else if n == 0 {
		z01.PrintRune('0')
		return
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
