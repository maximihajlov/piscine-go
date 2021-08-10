package Q9

func CountIf(f func(string) bool, tab []string) int {
	ans := 0
	for _, k := range tab {
		if f(k) {
			ans++
		}
	}
	return ans
}
