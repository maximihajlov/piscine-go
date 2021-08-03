package Q5

func ToUpper(s string) string {
	sArr := []rune(s)
	for i := range sArr {
		if 'a' <= sArr[i] && sArr[i] <= 'z' {
			sArr[i] -= 'a' - 'A'
		}
	}
	return string(sArr)
}
