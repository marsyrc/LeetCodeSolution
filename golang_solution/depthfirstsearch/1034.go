package depthfirstsearch

func colorBorder(grid [][]int, r0 int, c0 int, color int) [][]int {
	var dfs func(x, y int)
	m, n := len(grid), len(grid[0])
	inArea := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < m && y < n
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	flag := grid[r0][c0]
	dir := []int{0, -1, 0, 1, 0}
	dfs = func(x, y int) {
		visited[x][y] = true
		for i := 0; i < 4; i++ {
			newX, newY := x+dir[i], y+dir[i+1]
			if inArea(newX, newY) {
				if visited[newX][newY] {
					continue
				}
				if grid[newX][newY] != flag {
					grid[x][y] = color
				} else {
					visited[newX][newY] = true
					dfs(newX, newY)
				}
			} else {
				grid[x][y] = color
			}
		}
	}
	dfs(r0, c0)
	return grid
}
