package stack

import "math"

func largestRectangleArea(heights []int) int {
	n := len(heights)
	if n == 1 {
		return heights[0]
	}

	res := 0
	s := []int{}
	for i := 0; i < n; i++ {
		for len(s) != 0 && heights[s[len(s)-1]] > heights[i] {
			length := heights[s[len(s)-1]]
			s = s[:len(s)-1]
			weight := i
			if len(s) != 0 {
				weight = i - s[len(s)-1] - 1
			}
			res = max(res, weight*length)
		}
		s = append(s, i)
	}

	for len(s) != 0 {
		length := heights[s[len(s)-1]]
		s = s[:len(s)-1]
		weight := n
		if len(s) != 0 {
			weight = n - s[len(s)-1] - 1
		}
		res = max(res, length*weight)
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
