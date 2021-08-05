package piscine

func Split(s, sep string) []string {
	s = addSpaces(s, sep)
	return SplitWhiteSpaces(s)
}

func addSpaces(s string, sep string) string {
	sArr := []rune(s)
	for i := 0; i < len(s); i++ {
		if s[i] == sep[0] {
			isSep := true
			for j := range sep {
				if s[i+j] != sep[j] {
					isSep = false
					break
				}
			}
			if isSep {
				for k := i; k < i+len(sep); k++ {
					sArr[k] = ' '
				}
				i += len(sep)
			}
		}
	}
	return string(sArr)
}
