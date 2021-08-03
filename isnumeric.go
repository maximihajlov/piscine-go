package piscine

func IsNumeric(s string) bool {
	sArr := []rune(s)
	for i := range sArr {
		if !('0' <= sArr[i] && sArr[i] <= '9') {
			return false
		}
	}
	return true
}
