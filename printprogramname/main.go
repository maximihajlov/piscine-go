package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nArr := []rune(os.Args[0])
	for i := range nArr {
		z01.PrintRune(nArr[i])
	}
	z01.PrintRune('\n')
}
