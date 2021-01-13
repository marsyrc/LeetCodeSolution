package golang_alg

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
