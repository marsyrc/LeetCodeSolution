package unionfindset

import "fmt"

func maxNumEdgesToRemove(n int, edges [][]int) int {
	//greedy add edges
	res := 0
	uf := newUnionFindSet(n)
	//only common-used edges
	for _, v := range edges {
		if v[0] == 3 {
			if uf.isconnected(v[2]-1, v[1]-1) {
				res++
			} else {
				uf.union(v[2]-1, v[1]-1)
			}
		}
	}
	//copy into alice and bob 's
	ufa, ufb := copyUf(uf, n), copyUf(uf, n)
	//alice
	for _, v := range edges {
		if v[0] == 1 {
			if ufa.isconnected(v[2]-1, v[1]-1) {
				res++
			} else {
				ufa.union(v[2]-1, v[1]-1)
			}
		}
	}
	//bob
	for _, v := range edges {
		if v[0] == 2 {
			if ufb.isconnected(v[2]-1, v[1]-1) {
				res++
			} else {
				ufb.union(v[2]-1, v[1]-1)
			}
		}
	}
	fmt.Println(ufb.id)
	fmt.Println(ufa.id)
	if ufa.count != 1 || ufb.count != 1 {
		return -1
	}
	return res
}

type unionfind struct {
	id    []int
	size  []int
	count int
}

func copyUf(uf unionfind, n int) unionfind {
	newid := make([]int, n)
	newsz := make([]int, n)
	copy(newid, uf.id)
	copy(newsz, uf.size)
	nc := uf.count
	return unionfind{
		id:    newid,
		size:  newsz,
		count: nc,
	}
}
func newUnionFindSet(n int) unionfind {
	id := make([]int, n)
	size := make([]int, n)
	for i := range id {
		id[i] = i
	}
	for i := range size {
		size[i] = 1
	}
	return unionfind{
		id:    id,
		size:  size,
		count: n,
	}
}

func (uf *unionfind) find(p int) int {
	for uf.id[p] != p {
		p = uf.id[p]
	}
	return p
}

func (uf *unionfind) union(p, q int) {
	i, j := uf.find(p), uf.find(q)
	if i == j {
		return
	}
	if uf.size[i] > uf.size[j] {
		uf.id[j] = i
		uf.size[i] += uf.size[j]
	} else {
		uf.id[i] = j
		uf.size[j] += uf.size[i]
	}
	uf.count--
}

func (uf *unionfind) isconnected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}
