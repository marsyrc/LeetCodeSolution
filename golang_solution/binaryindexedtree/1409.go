package binaryindexedtree

func processQueries(queries []int, m int) []int {
	n := len(queries)
	bit := newBit(n + m)

	pos := make([]int, m+1)
	for i := 1; i <= m; i++ {
		pos[i] = n + i
		bit.update(n+i, 1)
	}

	ans := make([]int, n)
	for i := 0; i < n; i++ {
		cur := pos[queries[i]]
		bit.update(cur, -1)
		ans[i] = bit.query(cur)
		cur = n - i
		pos[queries[i]] = cur
		bit.update(cur, 1)
	}
	return ans
}

type BIT struct {
	arr []int
	n   int
}

func newBit(n int) BIT {
	return BIT{
		arr: make([]int, n+1),
		n:   n,
	}
}

func lowbit(x int) int {
	return x & (-x)
}

func (b *BIT) query(x int) int {
	res := 0
	for x != 0 {
		res += b.arr[x]
		x -= lowbit(x)
	}
	return res
}

func (b *BIT) update(x int, delta int) {
	for x <= b.n {
		b.arr[x] += delta
		x += lowbit(x)
	}
}
