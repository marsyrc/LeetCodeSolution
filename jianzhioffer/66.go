package jianzhioffer

func constructArr(a []int) []int {
	n := len(a)
	b := make([]int, n)
	if n == 0 {
		return []int{}
	}
	b[0] = 1
	for i := 1; i < n; i++ {
		b[i] = b[i-1] * a[i-1]
	}
	right := 1
	for i := n - 2; i >= 0; i-- {
		right *= a[i+1]
		b[i] *= right
	}
	return b
}
