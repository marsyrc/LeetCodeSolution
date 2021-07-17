package main

import "math"

func getMaxLen(nums []int) int {
	n := len(nums)
    neg, pos := make([]int, n + 1), make([]int, n + 1)
	res := 0
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			neg[i + 1] = 0
			pos[i + 1] = 0
		} else if nums[i] > 0 {
			if neg[i] == 0 {
				neg[i + 1] = 0
			} else {
				neg[i + 1] = neg[i] + 1
			}
			pos[i + 1] = pos[i] + 1
		} else {
			if neg[i] == 0 {
				pos[i + 1] = 0
			} else {
				pos[i + 1] = neg[i] + 1
			}
			neg[i + 1] = pos[i] + 1
		}

		res = max(res, pos[i + 1])
	}
	return res 
}

func max(a ...int) int {
	res := math.MinInt32
	for _, num := range a {
		if num > res {
			res = num
		}
	}
	return res 
}