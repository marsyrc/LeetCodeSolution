package floyd

import "math"

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt64
		}
	}
	for _, edge := range edges {
		u, v := edge[0], edge[1]
		dp[u][v] = edge[2]
		dp[v][u] = edge[2]
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if i == j || dp[i][k] == math.MaxInt64 || dp[k][j] == math.MaxInt64 {
					continue
				}
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k][j])
			}
		}
	}
	minNeighbors := math.MaxInt64
	res := -1
	for i := 0; i < n; i++ {
		cnt := 0
		for j := 0; j < n; j++ {
			if i != j && dp[i][j] <= distanceThreshold {
				cnt++
			}
		}
		if cnt <= minNeighbors {
			res = i
			minNeighbors = cnt
		}
	}
	return res
}

func min(a ...int) int {
	res := math.MaxInt64
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
