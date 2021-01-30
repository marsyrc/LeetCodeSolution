package twopointers

import "sort"

//较小的三数之和,n2复杂度
func threeSumSmaller(nums []int, target int) int {
	ans := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			cur := nums[i] + nums[left] + nums[right]
			if cur >= target {
				right--
			} else {
				ans += right - left
				left++
			}
		}
	}
	return ans
}
