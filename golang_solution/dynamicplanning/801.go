package dynamicplanning

func minSwap(A []int, B []int) int {
	n := len(A)
	dp := make([][]int, len(A))
	for i := range dp {
		dp[i] = make([]int, 2)
	}
	//dp[i][0] : not exchange cur idx; dp[i][1] : exchange
	dp[0][0] = 0
	dp[0][1] = 1
	for i := 1; i < n; i++ {
		if A[i] > A[i-1] && B[i] > B[i-1] {
			if B[i] > A[i-1] && A[i] > B[i-1] {
				dp[i][0] = min(dp[i-1][0], dp[i-1][1])
				dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
			} else {
				dp[i][0] = dp[i-1][0]
				dp[i][1] = dp[i-1][1] + 1
			}
		} else {
			dp[i][0] = dp[i-1][1]
			dp[i][1] = dp[i-1][0] + 1
		}
	}
	return min(dp[n-1][0], dp[n-1][1])
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
