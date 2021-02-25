package dynamicplanning

import "math"

func maximumScore(nums []int, mul []int) int {
	m := len(mul)
	n := len(nums)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	res := math.MinInt32
	for k := 1; k <= m; k++ {
		for i := 0; i <= k; i++ {
			if i == 0 {
				dp[i][k-i] = dp[i][k-i-1] + nums[n-k+i]*mul[k-1]
			} else if i == k {
				dp[i][k-i] = dp[i-1][k-i] + nums[i-1]*mul[k-1]
			} else {
				dp[i][k-i] = max(dp[i][k-i-1]+nums[n-k+i]*mul[k-1], dp[i-1][k-i]+nums[i-1]*mul[k-1])
			}
			if k == m {
				res = max(res, dp[i][k-i])
			}
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
