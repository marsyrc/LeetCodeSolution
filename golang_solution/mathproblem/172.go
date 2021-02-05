package mathproblem

//每隔 5 个数，出现一个 5，每隔 25 个数，出现 2 个 5，每隔 125 个数，出现 3 个 5
//ans = n / 5 + n / 25 + n / 125 + ...
func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		ans += n / 5
		n /= 5
	}
	return ans
}
