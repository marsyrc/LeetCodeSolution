package dynamicplanning

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 0 {
		return 0
	}

	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	if obstacleGrid[0][0] == 1 {
		return 0
	}
	dp[0][0] = 1
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] != 1 {
			dp[i][0] = 1
		} else {
			break
		}
	}

	for j := 1; j < n; j++ {
		if obstacleGrid[0][j] != 1 {
			dp[0][j] = 1
		} else {
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}
