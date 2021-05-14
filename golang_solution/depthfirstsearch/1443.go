package depthfirstsearch

func minTime(n int, edges [][]int, hasApple []bool) int {
	visited := make([]bool, n)
	graph := make(map[int][]int)
	for _, e := range edges {
		u, v := e[0], e[1]
		if _, ok := graph[u]; !ok {
			graph[u] = []int{}
		}
		if _, ok := graph[v]; !ok {
			graph[v] = []int{}
		}
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	var dfs func(cur int) int
	dfs = func(cur int) int {
		visited[cur] = true
		res := 0
		for _, neighbor := range graph[cur] {
			if visited[neighbor] {
				continue
			}
			cost := dfs(neighbor)
			if cost > 0 || hasApple[neighbor] {
				res += 2 + cost
			}
		}
		return res
	}
	return dfs(0)
}
