package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var nums [10]int
	tmp := n
	for tmp >= 1 {
		nums[tmp%10]++
		tmp /= 10
	}
	for i := '0'; i <= '9'; i++ {
		posi := int(i - '0')
		if nums[posi] >= 1 {
			for j := 0; j < nums[posi]; j++ {
				z01.PrintRune(i)
			}
		}
	}
}
