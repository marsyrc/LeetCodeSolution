package twopointers

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	ans := math.MaxInt32
	for i := 0; i < n; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if abs(ans, target) > abs(sum, target) {
				ans = sum
			}
			if sum == target {
				return target
			} else if sum < target {
				l++
				for l < r && nums[l] == nums[l-1] {
					l++
				}
			} else if sum > target {
				r--
				for l < r && nums[r] == nums[r+1] {
					r--
				}
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
