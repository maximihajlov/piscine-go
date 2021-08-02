package piscine

func Sqrt(nb int) int {
	i := 1
	for ; i*i < nb; i++ {
	}
	if i*i != nb {
		return 0
	}
	return i
}
