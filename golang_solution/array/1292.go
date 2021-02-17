package array

//二维矩阵前缀和
func maxSideLength(mat [][]int, threshold int) int {
	m, n := len(mat), len(mat[0])
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			dp[i][j] = mat[i-1][j-1] + dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1]
		}
	}
	ans := 0
	for k := 1; k <= min(m, n); k++ {
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				if i-k < 0 || j-k < 0 {
					continue
				}
				tmp := dp[i][j] - dp[i-k][j] - dp[i][j-k] + dp[i-k][j-k]
				if tmp <= threshold && k > ans {
					ans = k
				}
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
