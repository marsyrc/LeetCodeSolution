package array

func generateMatrix(n int) [][]int {
	top, bottom, left, right := 0, n-1, 0, n-1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	num := 0
	for top <= bottom && left <= right {
		for j := left; j <= right; j++ {
			num++
			res[top][j] = num
		}
		top++

		for i := top; i <= bottom; i++ {
			num++
			res[i][right] = num
		}
		right--

		for j := right; j >= left; j-- {
			num++
			res[bottom][j] = num
		}
		bottom--

		for i := bottom; i >= top; i-- {
			num++
			res[i][left] = num
		}
		left++
	}
	return res
}
