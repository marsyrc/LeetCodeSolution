package stringproblem

import "strings"

func myAtoi(s string) int {
	s = strings.TrimLeft(s, " ")
	if len(s) == 0 {
		return 0
	}
	symbol := 1
	x := 0
	if s[0] == '-' {
		symbol = -1
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	if len(s) > 0 && !isdigit(s[0]) {
		return 0
	}
	for i := 0; i < len(s) && isdigit(s[i]); i++ {
		x = 10*x + int(s[i]-'0')
		if !(x >= -(1<<31) && x <= 1<<31-1) {
			if symbol > 0 {
				return 1<<31 - 1
			} else {
				return -1 << 31
			}
		}
	}
	return x * symbol
}

func isdigit(c byte) bool {
	return c >= '0' && c <= '9'
}
