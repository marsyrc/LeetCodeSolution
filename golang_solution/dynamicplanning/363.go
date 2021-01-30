package dynamicplanning

import "math"

func maxSumSubmatrix(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	maxSum := math.MinInt32

	for l := 0; l < n; l++ {
		rowSum := make([]int, m)
		for r := l; r < n; r++ {
			for i := 0; i < m; i++ {
				rowSum[i] += matrix[i][r]
			}

			maxSum = max(maxSum, dpmax(rowSum, k))
			if maxSum == k {
				return k
			}
		}
	}
	return maxSum
}

func dpmax(arr []int, k int) int {
	max := math.MinInt32
	for l := 0; l < len(arr); l++ {
		sum := 0
		for r := l; r < len(arr); r++ {
			sum += arr[r]
			if sum > max && sum <= k {
				max = sum
			}
		}
	}
	return max
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
