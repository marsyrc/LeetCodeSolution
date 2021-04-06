package jianzhioffer

func exist(board [][]byte, word string) bool {
	dir := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	m := len(board)
	if m == 0 {
		return false
	}
	n := len(board[0])
	inArea := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < m && y < n
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	var check func(i, j, idx int) bool
	check = func(i, j, idx int) bool {
		if board[i][j] != word[idx] {
			return false
		}
		if idx == len(word)-1 {
			return true
		}
		visited[i][j] = true
		defer func() {
			visited[i][j] = false
		}()
		for _, d := range dir {
			nx, ny := i+d[0], j+d[1]
			if inArea(nx, ny) && !visited[nx][ny] {
				if check(nx, ny, idx+1) {
					return true
				}
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] && check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
