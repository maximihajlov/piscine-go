package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nArr := []rune(os.Args[0])
	for i := len(nArr) - 1; i >= 0; i-- {
		if nArr[i] == '/' {
			for j := i + 1; j < len(nArr); j++ {
				z01.PrintRune(nArr[j])
			}
		}
	}
	z01.PrintRune('\n')
}
