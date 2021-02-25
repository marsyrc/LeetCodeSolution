package topologicalsort

//floyd
func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for _, p := range prerequisites {
		dp[p[0]][p[1]] = true
	}
	for k := 0; k < n; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				if dp[i][k] && dp[k][j] {
					dp[i][j] = true
				}
			}
		}
	}
	res := make([]bool, len(queries))
	for i := range res {
		res[i] = dp[queries[i][0]][queries[i][1]]
	}
	return res
}
