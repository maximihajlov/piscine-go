package piscine

func IsAlpha(s string) bool {
	sArr := []rune(s)
	for i := range sArr {
		if !(('a' <= sArr[i] && sArr[i] <= 'z') || ('A' <= sArr[i] && sArr[i] <= 'Z') || ('0' <= sArr[i] && sArr[i] <= '9')) {
			return false
		}
	}
	return true
}
