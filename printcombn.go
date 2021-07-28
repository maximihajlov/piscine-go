package piscine

import (
	"github.com/01-edu/z01"
)

var nums = [10]rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func PrintCombN(size int) {
	data := []rune{}
	for i := 0; i < size; i++ {
		data = append(data, '0')
	}
	PrintCombNUtil(data, 0, 9, 0, size)
}

func PrintCombNUtil(data []rune, start int, end int, index int, size int) {
	if index == size {
		for j := 0; j < size; j++ {
			z01.PrintRune(data[j])
		}
		if !(data[size-1] == '9' && data[0] == rune('9'-size+1)) {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
		return
	}
	i := start
	for i <= end && end-i+1 >= size-index {
		data[index] = nums[i]
		PrintCombNUtil(data, i+1, end, index+1, size)
		i++
	}
}
