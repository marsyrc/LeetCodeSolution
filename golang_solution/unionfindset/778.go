package unionfindset

/**
利用sinked存储已经沉没的格子；
将格子按高度排序；
找到当前时间能淹没的格子，与已经淹没的邻居节点连通，这里重复查询很多，
优化复杂度常数可以使用优先队列的方法，但是懒得写了；
左上和右下连通了就返回时间。
*/
import "sort"

func swimInWater(grid [][]int) int {
	sinked := map[int]bool{}
	direc := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	n := len(grid)
	if n == 0 {
		return 0
	}
	inArea := func(x, y int) bool {
		return x >= 0 && x < n && y >= 0 && y < n
	}
	points := []point{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newPoint := point{
				i:      i,
				j:      j,
				height: grid[i][j],
			}
			points = append(points, newPoint)
		}
	}
	sort.Slice(points, func(i, j int) bool {
		return points[i].height < points[j].height
	})
	uf := newUnionFindSet(n * n)
	i := 0
	t := 0
	for !uf.isConnected(0, n*n-1) {
		for i < n*n && points[i].height == t {
			cur := points[i]
			id := cur.i*n + cur.j
			sinked[id] = true
			for _, d := range direc {
				nx, ny := cur.i+d[0], cur.j+d[1]
				_, ok := sinked[nx*n+ny]
				if inArea(nx, ny) && ok {
					uf.union(id, nx*n+ny)
				}
			}
			i++
		}
		t++
	}
	return t - 1
}

type point struct {
	i      int
	j      int
	height int
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
