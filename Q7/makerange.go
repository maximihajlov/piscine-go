package Q7

func MakeRange(min, max int) []int {
	var ans []int
	if max <= min {
		return ans
	}
	ans = make([]int, max-min)
	for i := range ans {
		ans[i] = min + i
	}
	return ans
}
