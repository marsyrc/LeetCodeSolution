package calculator

import "strconv"

//分治、递归
func diffWaysToCompute(input string) []int {
	tmp, err := strconv.Atoi(input)
	if err != nil {
		return []int{tmp}
	}
	ans := []int{}
	for idx, op := range input {
		if op == '+' || op == '-' || op == '/' || op == '*' {
			leftNums := diffWaysToCompute(input[:idx])
			rightNums := diffWaysToCompute(input[idx+1:])
			for _, l := range leftNums {
				for _, r := range rightNums {
					cur := 0
					if op == '+' {
						cur = l + r
					} else if op == '-' {
						cur = l - r
					} else if op == '*' {
						cur = l * r
					} else if op == '/' {
						cur = l / r
					}
					ans = append(ans, cur)
				}
			}
		}
	}
	return ans
}
