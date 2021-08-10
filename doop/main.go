package main

import (
	"fmt"
	"os"
	"strconv"
)

var (
	badArg = false
	maxPos = 9223372036854775807
	maxNeg = -9223372036854775808
)

func main() {
	if len(os.Args) == 4 {
		num1, err1 := strconv.Atoi(os.Args[1])
		num2, err2 := strconv.Atoi(os.Args[3])

		if err1 != nil || err2 != nil {
			return
		}

		switch os.Args[2] {
		case "+":
			sum(num1, num2)
		case "-":
			sum(num1, -num2)
		case "*":
			mult(num1, num2)
		case "/":
			if num2 != 0 {
				fmt.Println(num1 / num2)
			} else {
				fmt.Println("No division by 0")
			}
		case "%":
			if num2 != 0 {
				fmt.Println(num1 % num2)
			} else {
				fmt.Println("No modulo by 0")
			}
		}
	}
}

func sum(a, b int) {
	if !(b > 0 && a > maxPos-b || b <= 0 && a < maxNeg-b) {
		fmt.Println(a + b)
	}
}

func mult(a, b int) {
	if !(a == -1 && b == maxNeg || a == maxNeg && b == -1) {
		if a == 0 || b == 0 {
			fmt.Println("0")
		} else if (a > 0 && b > 0 || a < 0 && b < 0) && maxPos/b >= a {
			fmt.Println(a * b)
		} else if (a < 0 || b < 0) && maxNeg/b <= a {
			fmt.Println(a * b)
		}
	}
}
