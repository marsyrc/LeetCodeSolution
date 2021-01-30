package unionfindset

import "math"

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	uf := NewUnionFindSet(len(source))
	for _, v := range allowedSwaps {
		uf.Mix(v[0], v[1])
	}
	n := len(source)
	same := 0
	m := map[int]map[int]int{}
	for i := 0; i < n; i++ {
		fa := uf.Find(i)
		if _, ok := m[fa]; !ok {
			m[fa] = map[int]int{}
		}
		m[fa][source[i]]++
	}
	for i := 0; i < n; i++ {
		fa := uf.Find(i)
		if m[fa][target[i]] > 0 {
			same++
			m[fa][target[i]]--
		}
	}
	return n - same
}

type UnionFindSet struct {
	id    []int
	sz    []int
	count int
}

func NewUnionFindSet(n int) *UnionFindSet {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}

	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1
	}

	return &UnionFindSet{
		id:    id,
		sz:    sz,
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
	i := u.Find(p)
	j := u.Find(q)
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
	return u.Find(p) == u.Find(q)
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
