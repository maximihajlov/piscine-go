package piscine

func Join(strs []string, sep string) string {
	ans := ""
	for i := 0; i < len(strs); i++ {
		ans += strs[i]
		if i != len(strs)-1 {
			ans += sep
		}
	}
	return ans
}
