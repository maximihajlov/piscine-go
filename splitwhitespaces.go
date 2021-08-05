package piscine

func SplitWhiteSpaces(s string) []string {
	var ans []string
	for i := 0; i < len(s); i++ {
		if i == 0 || s[i-1] == ' ' || s[i-1] == '\n' || s[i-1] == '\t' {
			word := ""
			j := i
			for ; j < len(s) && s[j] != ' ' && s[j] != '\n' && s[j] != '\t'; j++ {
				word += string(s[j])
			}
			if word != "" {
				ans = append(ans, word)
			}
			i = j
		}
	}
	return ans
}
