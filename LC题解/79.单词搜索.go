/*
 * @lc app=leetcode.cn id=79 lang=golang
 *
 * [79] 单词搜索
 */

// @lc code=start
type pair struct {
	x, y int
}

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func exist(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])

	//visited 二位数组
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}

	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] {
			return false
		}
		if k == len(word)-1 {
			return true
		}

		vis[i][j] = true
		defer func() { vis[i][j] = false }() //回溯

		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}

	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// @lc code=end
