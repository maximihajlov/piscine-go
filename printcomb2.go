package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for i := '0'; i <= '9'; i++ {
		for j := '0'; j <= '9'; j++ {
			for k := i; k <= '9'; k++ {
				var tmp rune
				if int(k) > int(i) {
					tmp = '0'
				} else {
					tmp = j
				}
				for m := tmp; m <= '9'; m++ {
					if string(i)+string(j) != string(k)+string(m) {
						z01.PrintRune(i)
						z01.PrintRune(j)
						z01.PrintRune(' ')
						z01.PrintRune(k)
						z01.PrintRune(m)
						if string(i)+string(j)+string(k)+string(m) != "9899" {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
