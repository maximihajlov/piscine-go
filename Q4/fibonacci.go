package Q4

func Fibonacci(index int) int {
	if index < 2 {
		if index < 0 {
			return -1
		}
		return index
	}
	return Fibonacci(index-1) + Fibonacci(index-2)
}
