package unionfindset

func hitBricks(grid [][]int, hits [][]int) []int {
	//reset to 0 (for those hit)
	m, n := len(grid), len(grid[0])
	status := make([][]int, m)
	for i, row := range grid {
		status[i] = append([]int(nil), row...)
	}
	for _, h := range hits {
		i, j := h[0], h[1]
		status[i][j] = 0
	}

	//根据最终状态建立并查集
	//root : 虚拟的屋顶，
	//讨论root所在的分支的size变化，获取每一次增加砖块后变得稳定的砖块
	root := m * n
	uf := newUnionFindSet(root + 1)
	for i, row := range status {
		for j, v := range row {
			if v == 0 {
				continue
			}
			if i == 0 {
				uf.connect(i*n+j, root)
			}
			if i > 0 && status[i-1][j] == 1 {
				uf.connect(i*n+j, (i-1)*n+j)
			}
			if j > 0 && status[i][j-1] == 1 {
				uf.connect(i*n+j, i*n+j-1)
			}
		}
	}

	inArea := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < m && y < n
	}
	type pair struct{ x, y int }
	directions := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	ans := make([]int, len(hits))
	for i := len(hits) - 1; i >= 0; i-- {
		h := hits[i]
		r, c := h[0], h[1]
		if grid[r][c] == 0 {
			continue
		}
		preSize := uf.size[uf.find(root)]
		if r == 0 {
			uf.connect(r*n+c, root)
		}
		for _, d := range directions {
			newR, newC := r+d.x, c+d.y
			if inArea(newR, newC) && status[newR][newC] == 1 {
				uf.connect(r*n+c, newR*n+newC)
			}
		}
		cursize := uf.size[uf.find(root)]
		if cnt := cursize - preSize - 1; cnt > 0 {
			ans[i] = cnt
		}
		status[r][c] = 1
	}
	return ans
}

type unionFindSet struct {
	id    []int
	size  []int
	count int
}

func newUnionFindSet(n int) unionFindSet {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}

	size := make([]int, n)
	for i := range size {
		size[i] = 1
	}

	return unionFindSet{
		id:    id,
		size:  size,
		count: n,
	}
}

func (uf *unionFindSet) connect(p, q int) {
	i, j := uf.find(p), uf.find(q)
	if i == j {
		return
	}
	if uf.size[i] >= uf.size[j] {
		uf.id[j] = i
		uf.size[i] += uf.size[j]
	} else {
		uf.id[i] = j
		uf.size[j] += uf.size[i]
	}
	uf.count--
}

func (uf *unionFindSet) find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *unionFindSet) isConnected(p, q int) bool {
	if uf.find(p) == uf.find(q) {
		return true
	}
	return false
}
