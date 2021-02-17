package breadthfirstsearch

/*
	将n对情侣看成n个节点，有混乱配对看成节点之间连线
	则统计图中连通分量的大小，减去1，则是当前分量的交换次数
	利用visited避免重复访问，bfs方法实现
*/
func minSwapsCouples(row []int) int {
	m := len(row)
	n := m / 2
	graph := make([][]int, n)
	for i := 0; i < m; i += 2 {
		l, r := row[i]/2, row[i+1]/2
		graph[l] = append(graph[l], r)
		graph[r] = append(graph[r], l)
	}
	ans := 0
	visited := make([]bool, n)
	for i, vis := range visited {
		if vis {
			continue
		}
		visited[i] = true
		q := []int{i}
		cnt := 0
		for len(q) > 0 {
			cur := q[0]
			cnt++
			q = q[1:]
			for _, neighbor := range graph[cur] {
				if !visited[neighbor] {
					visited[neighbor] = true
					q = append(q, neighbor)
				}
			}
		}
		ans += cnt - 1
	}
	return ans
}

/*
	并查集方法
*/
func minSwapsCouples1(row []int) int {
	uf := newUnionFindSet(len(row) / 2)
	for i := 0; i < len(row); i += 2 {
		uf.union(row[i]/2, row[i+1]/2)
	}
	return len(row)/2 - uf.count
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
