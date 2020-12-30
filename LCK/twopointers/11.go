package twopointers

import "math"

func maxArea(height []int) int {
	n := len(height)
	l, r := 0, n-1
	res := -1
	for l < r {
		area := (r - l) * min(height[l], height[r])
		res = max(res, area)
		if height[l] > height[r] {
			r--
		} else {
			l++
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
