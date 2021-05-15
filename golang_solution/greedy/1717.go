package greedy

import "math"

func maximumGain(s string, x int, y int) int {
	stack := []byte{}
	s1, s2 := "", ""
	if x > y {
		s1 = "ab"
		s2 = "ba"
	} else {
		s2 = "ab"
		s1 = "ba"
	}
	x, y = max(x, y), min(x, y)
	res := 0
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])
		if len(stack) >= 2 && string(stack[len(stack)-2:]) == s1 {
			stack = stack[:len(stack)-2]
			res += x
		}
	}

	ls := []byte{}
	for i := 0; i < len(stack); i++ {
		ls = append(ls, stack[i])
		if len(ls) >= 2 && string(ls[len(ls)-2:]) == s2 {
			ls = ls[:len(ls)-2]
			res += y
		}
	}
	return res
}
func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
