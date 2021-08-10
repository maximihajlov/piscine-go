package Q9

func Any(f func(string) bool, a []string) bool {
	for _, k := range a {
		if f(k) {
			return true
		}
	}
	return false
}
