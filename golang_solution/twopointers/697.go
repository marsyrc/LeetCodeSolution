package twopointers

import "math"

func findShortestSubArray(nums []int) int {
	cnt := make(map[int]int)
	most := 1
	for _, n := range nums {
		cnt[n]++
		most = max(cnt[n], most)
	}
	cand := []int{}
	for i, v := range cnt {
		if v == most {
			cand = append(cand, i)
		}
	}
	res := len(nums)
	for _, v := range cand {
		l, r := 0, len(nums)-1
		for nums[l] != v {
			l++
		}
		for nums[r] != v {
			r--
		}
		res = min(res, r-l+1)
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
