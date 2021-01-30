package dynamicplanning

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	//boundary
	dp1 := nums[0]
	n := len(nums)
	if n == 1 {
		return dp1
	}
	dp2 := max(nums[0], nums[1])
	if n == 2 {
		return dp2
	}

	//dp[i] = max(dp[i - 1], dp[i - 2] + nums[i])
	for i := 2; i < n; i++ {
		tmp := dp2
		dp2 = max(dp1+nums[i], dp2)
		dp1 = tmp
	}
	return dp2
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
