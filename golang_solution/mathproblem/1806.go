package mathproblem

func reinitializePermutation(n int) int {
	res := 0
	t := 1
	if t < n/2 {
		t *= 2
	} else {
		t = (t-n/2)*2 + 1
	}
	res++
	for t != 1 {
		if t < n/2 {
			t *= 2
		} else {
			t = (t-n/2)*2 + 1
		}
		res++
	}
	return res
}
