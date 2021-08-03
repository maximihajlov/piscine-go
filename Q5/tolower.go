package Q5

func ToLower(s string) string {
	sArr := []rune(s)
	for i := range sArr {
		if 'A' <= sArr[i] && sArr[i] <= 'Z' {
			sArr[i] += 'a' - 'A'
		}
	}
	return string(sArr)
}
