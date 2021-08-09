package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	s := ""
	if len(os.Args) < 2 {
		z01.PrintRune('\n')
		return
	} else {
		for i := 1; i < len(os.Args); i++ {
			s += os.Args[i] + " "
		}
	}
	mirror(&s)
	PrintStr(s)
}

func mirror(s *string) {
	sArr := []rune(*s)
	j := len(sArr) - 1
	for i := 0; i < j; i++ {
		if isVowel(sArr[i]) {
			for ; j > i; j-- {
				if isVowel(sArr[j]) {
					sArr[i], sArr[j] = sArr[j], sArr[i]
					j--
					break
				}
			}
		}
	}
	*s = string(sArr)
}

func isVowel(l rune) bool {
	vowels := [...]rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

	for i := range vowels {
		if l == vowels[i] {
			return true
		}
	}
	return false
}

func PrintStr(s string) {
	sArr := []rune(s)
	for i := 0; i < len(s); i++ {
		z01.PrintRune(sArr[i])
	}
	z01.PrintRune('\n')
}
