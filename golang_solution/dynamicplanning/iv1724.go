package dynamicplanning

import (
	"math"
)

func getMaxMatrix(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, 4)
	maxSum := math.MinInt32

	for l := 0; l < n; l++ {
		rowSum := make([]int, m)
		for r := l; r < n; r++ {
			//updat rowSum from current [l, r]
			for i := 0; i < m; i++ {
				rowSum[i] += matrix[i][r]
			}
			sum := 0
			var up int
			for down, v := range rowSum {
				if sum < 0 {
					sum = 0
					up = down
				}
				sum += v
				if maxSum < sum {
					maxSum = sum
					res = []int{up, l, down, r}
				}
			}
		}
	}
	return res
}

func max(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if res < v {
			res = v
		}
	}
	return res
}
