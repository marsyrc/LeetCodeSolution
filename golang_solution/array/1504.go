package array

func numSubmat(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	row := make([][]int, m)
	for i := range row {
		row[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				row[i][j] = mat[i][j]
			} else if mat[i][j] != 0 {
				row[i][j] = 1 + row[i][j-1]
			}
		}
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			col := row[i][j]
			for k := i; k >= 0 && col != 0; k-- {
				col = min(col, row[k][j])
				ans += col
			}
		}
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
