package piscine

import (
	"github.com/01-edu/z01"
)

func PrintCombN(size int) {
	data := []rune{}
	for i := 0; i < size; i++ {
		data = append(data, '0')
	}
	PrintCombNUtil(data, 0, 9, 0, size)
	z01.PrintRune('\n')
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
	for i := start; i <= end && i <= end+index-size+1; i++ {
		data[index] = rune(i + '0')
		PrintCombNUtil(data, i+1, end, index+1, size)
	}
}
