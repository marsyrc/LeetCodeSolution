package hashmap

func isValidSudoku(board [][]byte) bool {
	col := make([]map[byte]bool, 9)
	row := make([]map[byte]bool, 9)
	box := make([]map[byte]bool, 9)
	for i := range col {
		col[i] = make(map[byte]bool)
	}
	for i := range row {
		row[i] = make(map[byte]bool)
	}
	for i := range col {
		box[i] = make(map[byte]bool)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			val := board[i][j]
			if col[j][val] || row[i][val] || box[(i/3)*3+j/3][val] {
				return false
			}
			col[j][val] = true
			row[i][val] = true
			box[i/3*3+j/3][val] = true
		}
	}
	return true
}
