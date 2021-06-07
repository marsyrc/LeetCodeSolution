package dynamicplanning

import "strings"

func findMaxForm(strs []string, m, n int) int {
	ls := len(strs)
	dp := make([][][]int, ls+1)
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, n+1)
		}
	}
	for i := 1; i <= ls; i++ {
		str := strs[i-1]
		cnt0, cnt1 := strings.Count(str, "0"), strings.Count(str, "1")
		for j := 0; j <= m; j++ {
			for k := 0; k <= n; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if j >= cnt0 && k >= cnt1 {
					dp[i][j][k] = max(dp[i][j][k], dp[i-1][j-cnt0][k-cnt1]+1)
				}
			}
		}
	}
	return dp[ls][m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
