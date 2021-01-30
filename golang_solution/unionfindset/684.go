package unionfindset

func findRedundantConnection(edges [][]int) []int {
	uf := NewUnionFindSet(len(edges) + 1)
	for i := range edges {
		u := edges[i][0]
		v := edges[i][1]
		if uf.Find(v) == uf.Find(u) {
			return edges[i]
		} else {
			uf.Mix(u, v)
		}
	}
	return []int{}
}

type UnionFindSet struct {
	id    []int
	count int
}

func NewUnionFindSet(n int) *UnionFindSet {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}

	return &UnionFindSet{
		id:    id,
		count: n,
	}
}

func (u *UnionFindSet) Find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *UnionFindSet) Mix(p, q int) {
	pRoot := u.Find(p)
	qRoot := u.Find(q)
	if pRoot == qRoot {
		return
	}
	u.id[pRoot] = qRoot
	u.count--
}

func (u *UnionFindSet) connected(p, q int) bool {
	return u.Find(p) == u.Find(q)
}
