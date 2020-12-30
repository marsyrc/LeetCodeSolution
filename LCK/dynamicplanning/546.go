package dynamicplanning

//hard
func removeBoxes(boxes []int) int {
	f := [100][100][100]int{}

	//refer the official code
	var helper func(l, r, k int) int
	helper = func(l, r, k int) int {
		if l > r {
			return 0
		}
		if f[l][r][k] != 0 {
			return f[l][r][k]
		}
		f[l][r][k] = helper(l, r-1, 0) + (k+1)*(k+1)
		for i := l; i < r; i++ {
			if boxes[i] == boxes[r] {
				f[l][r][k] = max(f[l][r][k], helper(l, i, k+1)+helper(i+1, r-1, 0))
			}
		}
		return f[l][r][k]
	}
	return helper(0, len(boxes)-1, 0)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
