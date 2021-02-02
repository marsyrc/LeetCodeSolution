package unionfindset

func numSimilarGroups(strs []string) int {
	n := len(strs)
	uf := newUnionFindSet(n)
	for i, s := range strs {
		for j := i + 1; j < n; j++ {
			t := strs[j]
			if !uf.isConnected(i, j) && check(s, t) {
				uf.union(i, j)
			}
		}
	}
	return uf.count
}

func check(s, t string) bool {
	cnt := 0
	if len(s) != len(t) {
		return false
	}
	for i := range s {
		if s[i] != t[i] {
			cnt++
			if cnt > 2 {
				return false
			}
		}
	}
	return true
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
