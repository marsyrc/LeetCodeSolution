package dynamicplanning

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	dp := func(nums []int) int {
		pre, cur := 0, 0
		for _, v := range nums {
			cur, pre = max(pre+v, cur), cur
		}
		return cur
	}
	s1 := nums[1:]
	s2 := nums[:len(nums)-1]
	return max(dp(s1), dp(s2))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
