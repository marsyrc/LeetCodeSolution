package depthfirstsearch

func updateBoard(board [][]byte, click []int) [][]byte {
	var dirX = []int{0, 1, 0, -1, 1, 1, -1, -1}
	var dirY = []int{1, 0, -1, 0, 1, -1, 1, -1}

	inArea := func(x, y int) bool {
			return x >= 0 && y >= 0 && x < len(board) && y < len(board[0])
	}

	var dfs func(x, y int)
	dfs = func(x, y int) {
			cnt := 0
			//统计地雷数量
			for i := 0; i < 8; i++ {
					nx, ny := x + dirX[i], y + dirY[i]
					if inArea(nx, ny) && board[nx][ny] == 'M' {
							cnt++
					}
			}
			
			if cnt > 0 {
					board[x][y] = byte(cnt + '0')
			} else {
					board[x][y] = 'B'
					for i := 0; i < 8; i++ {
							tx, ty := x + dirX[i], y + dirY[i]
							if inArea(tx, ty) && board[tx][ty] == 'E' {
									dfs(tx, ty)
							}
					}
			}
	}
	x, y := click[0], click[1]
	if board[x][y] == 'M' {
			board[x][y] = 'X'
	} else {
			dfs(x, y)
	}
	return board
}