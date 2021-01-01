package dynamicplanning

func findNumberOfLIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	count := make([]int, n)
	for i := range count {
		count[i] = 1
	}
	maxLen := 1
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j]
				} else if dp[i] == dp[j]+1 {
					count[i] += count[j]
				}

			}
		}
		maxLen = max(maxLen, dp[i])
	}

	res := 0
	for i := range count {
		if dp[i] == maxLen {
			res += count[i]
		}
	}
	return res
}
