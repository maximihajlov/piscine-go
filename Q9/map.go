package Q9

func Map(f func(int) bool, a []int) []bool {
	var ans []bool
	for _, k := range a {
		ans = append(ans, f(k))
	}
	return ans
}
