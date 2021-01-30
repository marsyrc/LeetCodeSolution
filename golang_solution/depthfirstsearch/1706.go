package depthfirstsearch

func findBall(grid [][]int) []int {
	if len(grid) == 0 {
		return []int{}
	}
	m, n := len(grid), len(grid[0])
	var dfs func(x int, y int) int
	dfs = func(x int, y int) int {
		if x == m {
			return y
		}
		if y > 0 && grid[x][y-1] == -1 && grid[x][y] == -1 {
			return dfs(x+1, y-1)
		}
		if y < n-1 && grid[x][y] == 1 && grid[x][y+1] == 1 {
			return dfs(x+1, y+1)
		}
		return -1
	}

	res := make([]int, len(grid[0]))
	for i := range res {
		res[i] = dfs(0, i)
	}
	return res
}
