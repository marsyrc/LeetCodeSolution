package breadthfirstsearch

func highestPeak(isWater [][]int) [][]int {
	m, n := len(isWater), len(isWater[0])
	inArea := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	q := []point{}
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isWater[i][j] == 1 {
				q = append(q, point{i, j})
			}
		}
	}
	dir := []int{0, -1, 0, 1, 0}
	H := 0
	for len(q) > 0 {
		sz := len(q)
		for _, node := range q {
			res[node.x][node.y] = H
		}
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			for i := 0; i < 4; i++ {
				nx, ny := cur.x+dir[i], cur.y+dir[i+1]
				if !inArea(nx, ny) || res[nx][ny] != -1 {
					continue
				}
				res[nx][ny] = -2
				q = append(q, point{nx, ny})
			}
		}
		H++
	}
	return res
}

type point struct {
	x int
	y int
}
