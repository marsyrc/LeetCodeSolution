package calculator

func calculate(s string) int {
	//在逻辑上去除括号
	stack := []int{}
	stack = append(stack, 1)

	res := 0
	num := 0
	op := 1

	for i := range s {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
			continue
		}

		res += op * num
		num = 0
		if c == '+' {
			op = stack[len(stack)-1]
		} else if c == '-' {
			op = -1 * stack[len(s)-1]
		} else if c == '(' {
			stack = append(stack, op)
		} else if c == ')' {
			stack = stack[:len(stack)-1]
		}
	}
	res += op * num
	return res
}
