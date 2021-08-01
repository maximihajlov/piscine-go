package piscine

func StrRev(s string) string {
	sArr := []rune(s)
	for i := 0; i < len(sArr)/2; i++ {
		sArr[i], sArr[-i-1] = sArr[-i-1], sArr[i]
	}
	return string(sArr)
}
