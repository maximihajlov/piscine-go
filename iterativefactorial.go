package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else {
		ans := 1
		for i := 1; i <= nb; i++ {
			ans *= i
		}
		return ans
	}
}
