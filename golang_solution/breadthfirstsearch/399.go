package breadthfirstsearch

type edge struct {
	to     int
	weight float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	id := map[string]int{}
	for _, v := range equations {
		if _, ok := id[v[0]]; !ok {
			id[v[0]] = len(id)
		}
		if _, ok := id[v[1]]; !ok {
			id[v[1]] = len(id)
		}
	}

	graph := make([][]edge, len(id))
	for i, eq := range equations {
		v, w := id[eq[0]], id[eq[1]]
		graph[v] = append(graph[v], edge{w, values[i]})
		graph[w] = append(graph[w], edge{v, 1 / values[i]})
	}

	bfs := func(start, end int) float64 {
		ratios := make([]float64, len(graph))
		ratios[start] = 1

		q := []int{start}
		for len(q) > 0 {
			sz := len(q)
			for i := 0; i < sz; i++ {
				v := q[0]
				q = q[1:]
				if v == end {
					return ratios[v]
				}
				for _, e := range graph[v] {
					if w := e.to; ratios[w] == 0 {
						ratios[w] = ratios[v] * e.weight
						q = append(q, w)
					}
				}
			}
		}
		return -1
	}

	ans := make([]float64, len(queries))
	for i, q := range queries {
		start, hasStart := id[q[0]]
		end, hasEnd := id[q[1]]
		if !hasStart || !hasEnd {
			ans[i] = -1
		} else {
			ans[i] = bfs(start, end)
		}
	}
	return ans
}
