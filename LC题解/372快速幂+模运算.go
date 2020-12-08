func superPow(a int, b []int) int {
	if len(b) == 0 {
		return 1
	}
	backend := b[len(b)-1]
	b = b[:len(b)-1]
	A := quickMul(a, backend)
	B := quickMul(superPow(a, b), 10)
	return (A * B) % 1337
}

func quickMul(x int, n int) int {
	if n == 0 {
		return 1
	}
	x %= 1337
	y := quickMul(x, n/2)
	if n%2 == 0 {
		return (y * y) % 1337
	}
	return (y * y * x) % 1337
}
