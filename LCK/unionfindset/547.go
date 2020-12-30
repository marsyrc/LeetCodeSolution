package unionfindset

func findCircleNum(M [][]int) int {
	if len(M) == 0 {
		return 0
	}
	friendUf := NewUnionFindSet(len(M))
	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M[0]); i++ {
			if M[i][j] == 1 && i != j {
				friendUf.Mix(i, j)
			}
		}
	}
	return friendUf.count
}

// UnionFindSet 并查集结构体
type UnionFindSet struct {
	id    []int // 分量id
	sz    []int //由触点索引的各个根节点对应的分量大小
	count int   // 连通分量数量
}

//NewUnionFindSet init a UF
func NewUnionFindSet(n int) *UnionFindSet {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}

	sz := make([]int, n)
	for i := range sz {
		sz[i] = i
	}

	return &UnionFindSet{
		id:    id,
		sz:    sz,
		count: n,
	}
}

// Find 查找根节点
func (u *UnionFindSet) Find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

// Mix 合并两个节点到同一个联通域
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
