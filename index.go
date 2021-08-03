package piscine

func Index(s string, toFind string) int {
	sArr := []rune(s)
	findArr := []rune(toFind)
	ans := -1
	for i := 0; i < len(sArr); i++ {
		ans = i
		for j := 0; j < len(findArr); j++ {
			if sArr[i+j] != findArr[j] {
				ans = -1
				break
			}
			if j == len(findArr)-1 {
				return ans
			}
		}
	}
	return ans
}
