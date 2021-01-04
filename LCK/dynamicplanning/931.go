package dynamicplanning

import "math"

func minFallingPathSum(A [][]int) int {
	n := len(A)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for j := 0; j < n; j++ {
		dp[n-1][j] = A[n-1][j]
	}

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < n; j++ {
			dp[i][j] = dp[i+1][j]
			if j > 0 {
				dp[i][j] = min(dp[i+1][j-1], dp[i][j])
			}
			if j < n-1 {
				dp[i][j] = min(dp[i+1][j+1], dp[i][j])
			}
			dp[i][j] += A[i][j]
		}
	}
	return min(dp[0]...)
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
