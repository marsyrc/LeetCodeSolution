package dynamicplanning

import "math"

func minCostII(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	n := len(costs)
	K := len(costs[0])

	//init
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, K)
	}
	for i := 1; i < n; i++ {
		for j := 0; j < K; j++ {
			dp[i][j] = math.MaxInt64
		}
	}
	for j := 0; j < K; j++ {
		dp[0][j] = costs[0][j]
	}

	for i := 1; i < n; i++ {
		for j := 0; j < K; j++ {
			for k := 0; k < K; k++ {
				if j == k {
					continue
				}
				dp[i][j] = min(dp[i-1][k]+costs[i][j], dp[i][j])
			}
		}
	}
	return min(dp[n-1]...)
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
