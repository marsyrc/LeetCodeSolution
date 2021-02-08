package greedy

//如果有一个堆的数量大于其余两个的和，那么无论如何三个堆都不可能为空。
//反之，三个堆都可以尽可能为空，如果总数为奇数，则三个堆总共剩一个，如果为偶数，则三个堆全为空。
func maximumScore(a int, b int, c int) int {
	if a+b <= c {
		return a + b
	} else if a+c <= b {
		return a + c
	} else if c+b <= a {
		return c + b
	} else {
		return (a + b + c) / 2
	}
}
