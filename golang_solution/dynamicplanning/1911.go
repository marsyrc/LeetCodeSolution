package dynamicplanning

import "math"

func maxAlternatingSum(nums []int) int64 {
	a, b := 0, math.MinInt64/2
	for _, v := range nums {
		a, b = max(a, b-v), max(b, a+v)
	}
	return int64(max(a, b))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
