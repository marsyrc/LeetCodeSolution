package jianzhioffer

import "math"

func dicesProbability(n int) []float64 {
	all := math.Pow(float64(6), float64(n))
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, 6*n+1)
	}
	for i := 1; i <= 6; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := i; j <= 6*i; j++ {
			for k := 1; k <= 6; k++ {
				if j-k <= 0 {
					break
				} else {
					dp[i][j] += dp[i-1][j-k]
				}
			}
		}
	}
	res := []float64{}
	for i := n; i <= 6*n; i++ {
		res = append(res, float64(dp[n][i])/all)
	}
	return res
}
