package binarysearch

//bs in 2dim array
//use i * n + j be the index of one element
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	l, h := 0, m*n-1
	for l <= h {
		mid := l + (h-l)/2
		midVal := matrix[mid/n][mid%n]
		if midVal == target {
			return true
		} else if midVal < target {
			l = mid + 1
		} else {
			h = mid - 1
		}
	}
	return false
}
