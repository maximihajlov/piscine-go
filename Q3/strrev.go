package piscine

func StrRev(s string) string {
	sArr := []rune(s)
	for i := 0; i < len(sArr)/2; i++ {
		tmp := sArr[i]
		sArr[i] = sArr[len(sArr)-i-1]
		sArr[len(sArr)-i-1] = tmp
	}
	s = ""
	for i := 0; i < len(sArr); i++ {
		s += string(sArr[i])
	}
	return s
}
