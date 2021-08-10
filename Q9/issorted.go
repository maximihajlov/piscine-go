package Q9

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) < 2 {
		return true
	}
	srtway := -1
	if f(a[0], a[1]) < 0 {
		srtway = 1
	}
	for i := range a[:len(a)-1] {
		if f(a[i], a[i+1])*srtway > 0 {
			return false
		}
	}
	return true
}
