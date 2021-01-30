package binarysearch

//快速幂
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1.0 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	y := myPow(x, n/2)
	if n%2 == 1 {
		return y * y * x
	}
	return y * y
}
