package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Print("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Print("Too many arguments")
		return
	}
	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(string(data))
}
