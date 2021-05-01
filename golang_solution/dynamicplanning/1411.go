package dynamicplanning

const mod = 1e9 + 7

func numOfWays(n int) int {
	types := []int{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				if i != j && j != k {
					types = append(types, i*9+j*3+k)
				}
			}
		}
	}

	typesCnt := len(types)
	related := make([][]int, typesCnt)
	for i := range related {
		related[i] = make([]int, typesCnt)
	}

	for i := 0; i < typesCnt; i++ {
		x1, x2, x3 := types[i]/9, types[i]/3%3, types[i]%3
		for j := 0; j < typesCnt; j++ {
			y1, y2, y3 := types[j]/9, types[j]/3%3, types[j]%3
			if x1 != y1 && x2 != y2 && x3 != y3 {
				related[i][j] = 1
			}
		}
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, typesCnt)
	}
	for i := 0; i < typesCnt; i++ {
		dp[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		for j := 0; j < typesCnt; j++ {
			for k := 0; k < typesCnt; k++ {
				if related[k][j] == 1 {
					dp[i][j] += dp[i-1][k]
					dp[i][j] %= mod
				}
			}
		}
	}

	ans := 0
	for i := 0; i < typesCnt; i++ {
		ans += dp[n][i]
		ans %= mod
	}
	return ans
}
