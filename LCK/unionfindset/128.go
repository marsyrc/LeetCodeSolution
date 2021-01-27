package unionfindset

func longestConsecutive(nums []int) int {
	//key : nums[i]; value : i
	uf := NewUnionFindSet(len(nums))
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		cur := nums[i]
		if _, ok := m[cur]; ok {
			continue
		}
		m[cur] = i
		if idx, ok := m[cur-1]; ok {
			uf.Mix(i, idx)
		}
		if idx, ok := m[cur+1]; ok {
			uf.Mix(i, idx)
		}
	}
	res := 0
	for _, v := range uf.sz {
		if v > res {
			res = v
		}
	}
	return res
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
