package jianzhioffer

func maxSubArray(nums []int) int {
	res := nums[0]
	cur := nums[0]
	for i := 1; i < len(nums); i++ {
		if cur > 0 {
			cur += nums[i]
		} else {
			cur = nums[i]
		}
		res = max(res, cur)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
