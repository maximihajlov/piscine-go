package Q9

func ForEach(f func(int), a []int) {
	for _, k := range a {
		f(k)
	}
}
