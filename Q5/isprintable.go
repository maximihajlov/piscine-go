package Q5

func IsPrintable(s string) bool {
	sArr := []rune(s)
	for i := range sArr {
		if !(' ' <= sArr[i] && sArr[i] <= 127) {
			return false
		}
	}
	return true
}
