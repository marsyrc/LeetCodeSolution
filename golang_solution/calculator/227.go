package calculator

func calculate(s string) int {
	stack := make([]int, 0, len(s))
	operator := make([]int, 0, len(s))
	num := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num = 0
			for i < len(s) && s[i] <= '9' && s[i] >= '0' {
				num = num*10 + int(s[i]-'0')
				i++
			}
			if len(operator) > 0 && operator[len(operator)-1] > 1 {
				if operator[len(operator)-1] == 2 {
					stack[len(stack)-1] *= num
				} else {
					stack[len(stack)-1] /= num
				}
				operator = operator[:len(operator)-1]
			} else {
				stack = append(stack, num)
			}
			i--
		case '+':
			operator = append(operator, 1)
		case '-':
			operator = append(operator, -1)
		case '*':
			operator = append(operator, 2)
		case '/':
			operator = append(operator, 3)
		default:
		}
	}
	for len(operator) > 0 {
		stack[1] = stack[0] + operator[0]*stack[1]
		operator = operator[1:]
		stack = stack[1:]
	}
	return stack[0]
}
