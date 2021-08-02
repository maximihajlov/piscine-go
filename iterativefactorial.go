package piscine

func IterativeFactorial(nb int) int {
	ans := 1
	for i := 1; i <= nb; i++ {
		ans *= i
	}
	return ans
}
