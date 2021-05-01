package BIT

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
