package array

func maxRotateFunction(A []int) int {

	total, ans := 0, 0
	for i, a := range A {
		total += a
		ans += i * a
	}
	cur, n := ans, len(A)
	for i := range A {
		cur += total - n*A[n-1-i]
		if cur > ans {
			ans = cur
		}
	}

	return ans
}
