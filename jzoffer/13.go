package jianzhioffer

//bfs
type pair struct {
	x, y int
}

var directions = []pair{{1, 0}, {0, 1}}

func movingCount(m int, n int, k int) int {
	digitSum := func(x int) (res int) {
		for ; x > 0; x = x / 10 {
			res += x % 10
		}
		return
	}
	var check func(i, j int) bool
	check = func(i, j int) bool {
		sum := digitSum(i) + digitSum(j)
		if sum <= k {
			return true
		}
		return false
	}
	var queue = []pair{{0, 0}}
	counts := 1
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := size; i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				x := node.x + dir.x
				y := node.y + dir.y
				if x < m && y < n && x >= 0 && y >= 0 && !visited[x][y] {
					if !check(x, y) {
						continue
					} else {
						queue = append(queue, pair{x, y})
						visited[x][y] = true
						counts++
					}
				} else {
					continue
				}
			}
		}
	}
	return counts
}
