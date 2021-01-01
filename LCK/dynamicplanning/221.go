package dynamicplanning

import "math"

func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	l := 0
	//init from matrix
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = 1
				l = 1
			}
		}
	}

	//consider the left\up\leftup
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
				l = max(l, dp[i][j])
			}
		}
	}
	return l * l
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

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
