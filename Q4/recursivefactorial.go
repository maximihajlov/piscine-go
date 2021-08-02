package Q4

func RecursiveFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	} else if nb > 1 {
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 1
	}
}
