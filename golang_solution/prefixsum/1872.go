package prefixsum

func stoneGameVIII(a []int) int {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	f := sum[n]
	for i := n - 1; i > 1; i-- {
		f = max(f, sum[i]-f)
	}
	return f
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
