package stack

import (
	"strconv"
	"strings"
)

func exclusiveTime(n int, logs []string) []int {
	res := make([]int, n)
	lastIdx := -1
	stack := []int{}
	for _, log := range logs {
		strs := strings.Split(log, ":")
		cur, _ := strconv.Atoi(strs[0])
		time, _ := strconv.Atoi(strs[2])
		if strs[1] == "start" {
			if len(stack) > 0 {
				res[stack[len(stack)-1]] += time - lastIdx
			}
			stack = append(stack, cur)
			lastIdx = time
		} else {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res[top] += time - lastIdx + 1
			lastIdx = time + 1
		}
	}
	return res
}
