package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if isBigger(a[i], a[j]) {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func isBigger(s1, s2 string) bool {
	if s2 == "" {
		return true
	}
	s1Arr := []rune(s1)
	s2Arr := []rune(s2)
	for i := 0; i < len(s1Arr) && i < len(s2Arr); i++ {
		if s1Arr[i] > s2Arr[i] {
			return true
		} else if s1Arr[i] < s2Arr[i] {
			return false
		}
	}
	return false
}
