package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for i := 1; i < len(args); i++ {
		for j := i; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
	for i := 0; i < len(args); i++ {
		arg := []rune(args[i])
		for j := range arg {
			z01.PrintRune(arg[j])
		}
		z01.PrintRune('\n')
	}
}
