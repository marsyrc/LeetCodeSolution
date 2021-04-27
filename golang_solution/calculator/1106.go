package calculator

func parseBoolExpr(expression string) bool {
	start, end := 2, len(expression)-1
	switch expression[0] {
	case '!':
		return Not(expression[start:end])
	case '&':
		return AndOr(expression[start:end], true)
	case '|':
		return AndOr(expression[start:end], false)
	default:
		return expression == "t"
	}
}

// flag &: true |: false
func AndOr(exp string, flag bool) bool {
	pre, idx := 0, 0
	for i := 0; i < len(exp); i++ {
		if exp[i] == '(' {
			if pre == 0 {
				idx = i
			}
			pre++
		} else if exp[i] == ')' {
			pre--
			if pre == 0 {
				t := parseBoolExpr(exp[idx-1 : i+1])
				if !t && flag {
					return false
				}
				if t && !flag {
					return true
				}
			}
		} else {
			if pre <= 0 {
				if exp[i] == 'f' && flag {
					return false
				}
				if exp[i] == 't' && !flag {
					return true
				}
			}
		}
	}
	return flag
}

func Not(exp string) bool {
	if len(exp) == 1 {
		return !(exp == "t")
	}
	return !parseBoolExpr(exp)
}
