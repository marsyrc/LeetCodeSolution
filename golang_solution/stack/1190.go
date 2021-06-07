package stack

func reverseParentheses(s string) string {
	i := 0
	var recur func() string
	recur = func() string {
		str := ""
		for i < len(s) {
			if s[i] == '(' {
				i++
				str += recur()
			} else if s[i] == ')' {
				return reverse(str)
			} else {
				str += string(s[i])
			}
			i++
		}
		return str
	}
	return recur()
}

func reverse(a string) string {
	n := len(a)
	as := []byte(a)
	i, j := 0, n-1
	for i < j {
		as[i], as[j] = as[j], as[i]
		i++
		j--
	}
	return string(as)
}
