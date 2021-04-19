package stack

import "strings"

func lengthLongestPath(input string) int {
	stack := []int{0}
	ans := 0
	strs := strings.Split(input, "\n")
	for _, s := range strs {
		level := strings.LastIndex(s, "\t") + 1
		for level+1 < len(stack) {
			stack = stack[:len(stack)-1]
		}
		length := stack[len(stack)-1] + len(s) - level + 1
		stack = append(stack, length)
		if strings.Contains(s, ".") {
			ans = max(ans, length-1)
		}
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
