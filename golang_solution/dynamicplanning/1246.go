package dynamicplanning

//区间dp
func minimumMoves(arr []int) int {
	n := len(arr)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	//l = 1
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	//l = 2
	for i := 0; i < n-1; i++ {
		if arr[i] == arr[i+1] {
			dp[i][i+1] = 1
		} else {
			dp[i][i+1] = 2
		}
	}
	for l := 3; l <= n; l++ {
		for i := 0; i+l-1 < n; i++ {
			j := i + l - 1
			dp[i][j] = l
			for k := i; k < j; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j])
			}
			if arr[i] == arr[j] {
				dp[i][j] = min(dp[i][j], dp[i+1][j-1])
			}
		}
	}
	return dp[0][n-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
