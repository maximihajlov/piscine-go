package Q4

func IterativeFactorial(nb int) int {
	if nb < 0 || nb > 20 {
		return 0
	} else {
		ans := 1
		for i := 1; i <= nb; i++ {
			ans *= i
		}
		return ans
	}
}
