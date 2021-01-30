package dynamicplanning

import "math"

//same as 64.go
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func minimumTotal(triangle [][]int) (res int) {
	if len(triangle) == 0 {
		return 0
	}

	m := len(triangle)
	f := make([][]int, m)
	for i := range f {
		f[i] = make([]int, i+1)
	}

	f[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		f[i][0] = f[i-1][0] + triangle[i][0]
		for j := 1; j < i; j++ {
			f[i][j] = min(f[i-1][j], f[i-1][j-1]) + triangle[i][j]
		}
		f[i][i] = f[i-1][i-1] + triangle[i][i]
	}
	res = math.MaxInt32
	for j := 0; j < m; j++ {
		res = min(res, f[m-1][j])
	}
	return res
}
