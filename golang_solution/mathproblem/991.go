package mathproblem

func brokenCalc(x int, y int) int {
	res := 0
	for y > x {
		if  y & 1 == 1 {
			y++
		} else {
			y >>= 1
		}
		res++
	}
	return res + (x - y)
}