package mathproblem

func clumsy(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	if N == 2 {
		return 2 * 1
	}
	if N == 3 {
		return 3 * 2 / 1
	}
	k := N / 4
	remainder := N % 4
	res := N*(N-1)/(N-2) + (N - 3)
	N -= 4
	for i := 1; i < k; i++ {
		res -= N * (N - 1) / (N - 2)
		res += (N - 3)
		N -= 4
	}
	res -= clumsy(remainder)
	return res
}
