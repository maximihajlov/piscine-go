package piscine

func FindNextPrime(nb int) int {
	for i := nb; ; i++ {
		if IsPrime(i) {
			return i
		}
	}
}

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
