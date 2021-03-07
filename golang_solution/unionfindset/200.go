package unionfindset

func numIslands(grid [][]byte) int {
	m, n := len(grid), len(grid[0])
	uf := newUnionFindSet(n * m)
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	inArea := func(i, j int) bool {
		return i >= 0 && j >= 0 && j < n && i < m
	}
	space := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '0' {
				space++
				continue
			}
			for _, d := range dir {
				if inArea(i+d[0], j+d[1]) && grid[i+d[0]][j+d[1]] == '1' {
					uf.union(i*n+j, (i+d[0])*n+j+d[1])
				}
			}
		}
	}
	return uf.count - space
}

type UnionFindSet struct {
	id    []int
	sz    []int
	count int
}

func newUnionFindSet(n int) UnionFindSet {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}

	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1
	}

	return UnionFindSet{
		id:    id,
		sz:    sz,
		count: n,
	}
}

func (u *UnionFindSet) find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *UnionFindSet) union(p, q int) {
	i := u.find(p)
	j := u.find(q)
	if i == j {
		return
	}
	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}
	u.count--
}

func (u *UnionFindSet) isConnected(p, q int) bool {
	return u.find(q) == u.find(p)
}
