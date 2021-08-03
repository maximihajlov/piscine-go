package Q5

func IsUpper(s string) bool {
	sArr := []rune(s)
	for i := range sArr {
		if !('A' <= sArr[i] && sArr[i] <= 'Z') {
			return false
		}
	}
	return true
}
