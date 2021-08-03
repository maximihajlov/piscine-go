package piscine

func IsLower(s string) bool {
	sArr := []rune(s)
	for i := range sArr {
		if !('a' <= sArr[i] && sArr[i] <= 'z') {
			return false
		}
	}
	return true
}
