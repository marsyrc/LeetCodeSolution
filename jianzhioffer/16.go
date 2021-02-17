package jianzhioffer

func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / quickMultify(x, -n)
	}
	return quickMultify(x, n)
}

func quickMultify(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	y := quickMultify(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}
