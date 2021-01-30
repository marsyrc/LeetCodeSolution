package mathproblem

func reverse(x int) int {
	y := 0
	for x != 0 {
		y = 10*y + x%10
		if !(y >= -(1<<31) && y <= (1<<31)-1) {
			return 0
		}
		x /= 10
	}
	return y
}
