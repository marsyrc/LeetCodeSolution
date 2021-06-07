package depthfirstsearch

func minimumJumps(forbidden []int, a int, b int, x int) int {
	visited := make([]bool, 8000)
	for _, f := range forbidden {
		visited[f] = true
	}
	res := -1
	var dfs func(cur int, step int, fromBack bool)
	dfs = func(cur int, step int, fromBack bool) {
		if !(res == -1 && cur >= 0 && cur <= 6000) {
			return
		}
		if cur == x {
			res = step
			return
		}
		if !visited[cur+a] {
			visited[cur+a] = true
			dfs(cur+a, step+1, false)
		}
		if !fromBack && cur-b > 0 && !visited[cur-b] {
			dfs(cur-b, step+1, true)
		}
	}
	dfs(0, 0, false)
	return res
}
