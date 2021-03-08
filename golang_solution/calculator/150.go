package calculator

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		if s, err := strconv.Atoi(v); err == nil {
			stack = append(stack, s)
		} else {
			a := stack[len(stack)-1]
			b := stack[len(stack)-2]

			if v == "+" {
				newv := a + b
				stack = append(stack[:len(stack)-2], newv)
			} else if v == "-" {
				newv := b - a
				stack = append(stack[:len(stack)-2], newv)
			} else if v == "*" {
				newv := a * b
				stack = append(stack[:len(stack)-2], newv)
			} else if v == "/" {
				newv := b / a
				stack = append(stack[:len(stack)-2], newv)
			} else {
				return 0
			}
		}
	}
	return stack[0]
}
