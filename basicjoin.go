package piscine

func BasicJoin(elems []string) string {
	ans := ""
	for i := 0; i < len(elems); i++ {
		ans += elems[i]
	}
	return ans
}
