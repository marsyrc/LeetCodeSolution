package array

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])
	firstColContainsZero := false
	firstRowContainsZero := false
	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowContainsZero = true
			break
		}
	}
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColContainsZero = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstRowContainsZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}

	if firstColContainsZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}
}
