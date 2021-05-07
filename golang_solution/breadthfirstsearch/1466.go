package breadthfirstsearch

type node struct {
	des int
	dir bool //src -> des : true
}

func minReorder(n int, connections [][]int) int {
	res := 0
	graph := make(map[int][]node)
	q := []int{0}
	//bulid map
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		if _, ok := graph[u]; !ok {
			graph[u] = []node{}
		}
		if _, ok := graph[v]; !ok {
			graph[v] = []node{}
		}
		graph[u] = append(graph[u], node{v, true})
		graph[v] = append(graph[v], node{u, false})
	}
	visited := make([]bool, n)
	visited[0] = true

	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			for _, nei := range graph[cur] {
				if !visited[nei.des] {
					visited[nei.des] = true
					q = append(q, nei.des)
					if nei.dir == true {
						res++
					}
				}
			}
		}
	}
	return res
}
