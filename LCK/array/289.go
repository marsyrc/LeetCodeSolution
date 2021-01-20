package array

//复合状态解决
func gameOfLife(board [][]int) {
	if len(board) == 0 {
		return
	}
	m, n := len(board), len(board[0])
	diretions := [][]int{{0, 1}, {0, -1}, {1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, 0}, {-1, -1}}
	inArea := func(x, y int) bool {
		return x >= 0 && y >= 0 && x < m && y < n
	}
	//2 : dead to alive; -1 : alive to dead
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			cnt := 0
			for _, d := range diretions {
				x, y := i+d[0], j+d[1]
				if !inArea(x, y) {
					continue
				}
				if board[x][y] == 1 || board[x][y] == -1 {
					cnt++
				}
			}
			if board[i][j] == 1 {
				if cnt < 2 || cnt > 3 {
					board[i][j] = -1
				}
			} else if board[i][j] == 0 {
				if cnt == 3 {
					board[i][j] = 2
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] > 0 {
				board[i][j] = 1
			} else {
				board[i][j] = 0
			}
		}
	}
}
