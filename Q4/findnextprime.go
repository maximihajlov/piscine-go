package Q4

func FindNextPrime(nb int) int {
	for i := nb; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
}
