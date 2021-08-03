package Q5

func AlphaCount(s string) int {
	sArr := []rune(s)
	count := 0
	for i := range sArr {
		l := sArr[i]
		if 'A' <= l && l <= 'Z' || 'a' <= l && l <= 'z' {
			count++
		}
	}
	return count
}
