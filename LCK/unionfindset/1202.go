package unionfindset

import "sort"

// 利用并查集记录索引之间的联通度
// 利用map记录每一个根节点下有什么字母和字母计数
// 在根结点下用字典计数，切片排序并计数，小顶堆都很麻烦而且35/36，直接[]byte记录所有的字符串再按字母序排序就好了
// 用掉一个字符，就更新当前father下的[]byte
// 用了quick-find\quick-union并查集时间都不太理想，还得路径压缩+加权quick-union
func smallestStringWithSwaps(s string, pairs [][]int) string {
	uf := NewUnionFindSet(len(s))
	if len(s) == 0 {
		return ""
	}
	for _, p := range pairs {
		uf.Mix(p[0], p[1])
	}

	m := map[int][]byte{}
	for i := range s {
		father := uf.Find(i)
		if _, ok := m[father]; !ok {
			m[father] = []byte{}
		}
		m[father] = append(m[father], s[i])
	}

	for fa := range m {
		sort.Slice(m[fa], func(i, j int) bool {
			return m[fa][i] < m[fa][j]
		})
	}

	t := []byte{}
	for i := range s {
		father := uf.Find(i)
		t = append(t, m[father][0])
		m[father] = m[father][1:]
	}
	return string(t)
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

// Find 查找根节点
func (u *UnionFindSet) Find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

// Union 合并两个节点到同一个联通域
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
