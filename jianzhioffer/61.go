package jianzhioffer

import "math"

func isStraight(nums []int) bool {
	seen := map[int]int{}
	maxN, minN := -1, 14
	for _, n := range nums {
		if n == 0 {
			continue
		}
		minN = min(minN, n)
		maxN = max(maxN, n)
		if _, ok := seen[n]; ok {
			return false
		}
		seen[n]++
	}
	return maxN-minN < 5
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
