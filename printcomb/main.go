package main

import "github.com/01-edu/z01"

func main() {
	for i := '0'; i <= '9'; i++ {
		for j := i + 1; j <= '9'; j++ {
			for k := j + 1; k <= '9'; k++ {
				z01.PrintRune(i)
				z01.PrintRune(j)
				z01.PrintRune(k)

				if string(i)+string(j)+string(k) != "789" {
					z01.PrintRune(',')
					z01.PrintRune(' ')
				}
			}
		}
	}
	z01.PrintRune('\n')
}
