package Q7

func ConcatParams(args []string) string {
	ans := ""
	for i := range args {
		ans += args[i]
		if i != len(args)-1 {
			ans += "\n"
		}
	}
	return ans
}
