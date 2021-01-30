package topologicalsort

func findOrder(numCourses int, prerequisites [][]int) []int {
	var result []int

	edges := make([][]int, numCourses)
	indegree := make([]int, numCourses)
	for _, v := range prerequisites {
		edges[v[1]] = append(edges[v[1]], v[0])
		indegree[v[0]]++
	}

	q := []int{}
	for i := 0; i < numCourses; i++ {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}

	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		result = append(result, cur)

		for _, v := range edges[cur] {
			indegree[v]--
			if indegree[v] == 0 {
				q = append(q, v)
			}
		}
	}

	if len(result) == numCourses {
		return result
	}
	return []int{}
}
