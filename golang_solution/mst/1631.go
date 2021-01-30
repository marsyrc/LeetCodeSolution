package mst

import "sort"

func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	edges := []edge{}
	//build edges
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			loc := n*i + j
			if i > 0 {
				newedge := edge{
					a:      loc - n,
					b:      loc,
					weight: abs(heights[i][j] - heights[i-1][j]),
				}
				edges = append(edges, newedge)
			}
			if j > 0 {
				newedge := edge{
					a:      loc - 1,
					b:      loc,
					weight: abs(heights[i][j] - heights[i][j-1]),
				}
				edges = append(edges, newedge)
			}
		}
	}
	//kruskal
	uf := newUnionFindSet(m * n)
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})
	for _, e := range edges {
		uf.union(e.a, e.b)
		if uf.isConnected(0, m*n-1) {
			return e.weight
		}
	}
	return -1
}

type edge struct {
	a      int
	b      int
	weight int
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
