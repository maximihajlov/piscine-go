package Q5

func Capitalize(s string) string {
	sArr := []rune(s)
	if 'a' <= sArr[0] && sArr[0] <= 'z' {
		sArr[0] += 'A' - 'a'
	}
	for i := 1; i < len(sArr); i++ {
		if !('0' <= sArr[i-1] && sArr[i-1] <= '9' || 'a' <= sArr[i-1] && sArr[i-1] <= 'z' || 'A' <= sArr[i-1] && sArr[i-1] <= 'Z') {
			if 'a' <= sArr[i] && sArr[i] <= 'z' {
				sArr[i] += 'A' - 'a'
			}
		} else if 'A' <= sArr[i] && sArr[i] <= 'Z' {
			sArr[i] -= 'A' - 'a'
		}
	}
	return string(sArr)
}
