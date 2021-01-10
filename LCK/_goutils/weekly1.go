package _goutils

import "math"

func minimumHammingDistance(source []int, target []int, allowedSwaps [][]int) int {
	uf := NewUnionFindSet(len(source))
	for _, v := range allowedSwaps {
		uf.Mix(v[0], v[1])
	}
	n := len(source)
	ans := 0
	m := make(map[int][]int)
	for i := 0; i < n; i++ {
		m[target[i]] = make([]int, 0)
		m[target[i]] = append(m[target[i]], i)
	}

	for i := 0; i < n; i++ {
		key := source[i]
		flag := false
		for index, j := range m[key] {
			if uf.connected(i, j) {
				m[key] = append(m[key][:index], m[key][index+1:]...)
				flag = true
				break
			}
		}
		if !flag {
			ans++
		}
	}
	return ans
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

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
