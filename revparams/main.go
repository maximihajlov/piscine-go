package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for i := len(os.Args) - 1; i > 0; i-- {
		arg := []rune(os.Args[i])
		for j := range arg {
			z01.PrintRune(arg[j])
		}
		z01.PrintRune('\n')
	}
}
