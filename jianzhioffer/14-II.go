package jianzhioffer

func cuttingRope(n int) int {
	if n <= 3 {
		return n - 1
	}
	const mod = 1e9 + 7
	res := 1
	for n > 4 {
		n -= 3
		res = res * 3 % mod
	}
	return res * n % mod
}
