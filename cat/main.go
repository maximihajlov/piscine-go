package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		stdUtil()
	} else if len(os.Args) == 2 {
		readFile(os.Args[1])
	} else {
		for _, arg := range os.Args[1:] {
			readFile(arg)
		}
	}
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		printStr("ERROR: " + err.Error() + "\n")
		os.Exit(1)
	}
	printStr(string(data))
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func stdUtil() {
	printStr("Hello\n")
}
