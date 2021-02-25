package dynamicplanning

import "math"

func maxDotProduct(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	res := math.MinInt32
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i][j] = nums1[i] * nums2[j]
			if i > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j])
			}
			if j > 0 {
				dp[i][j] = max(dp[i][j], dp[i][j-1])
			}
			if i > 0 && j > 0 {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+nums1[i]*nums2[j])
			}
			res = max(res, dp[i][j])
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
