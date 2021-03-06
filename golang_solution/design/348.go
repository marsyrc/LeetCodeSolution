package design

import "math"

type TicTacToe struct {
	N    int
	path []int
}

/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	t := TicTacToe{
		N:    n,
		path: make([]int, n*2+2),
	}
	return t
}

/** Player {player} makes a move at ({row}, {col}).
  @param row The row of the board.
  @param col The column of the board.
  @param player The player, can be either 1 or 2.
  @return The current winning condition, can be either:
          0: No one wins.
          1: Player 1 wins.
          2: Player 2 wins. */
func (this *TicTacToe) Move(row int, col int, player int) int {
	add := int(math.Pow(-1, float64(player)))
	this.path[row] += add
	this.path[this.N+col] += add
	if row == col {
		this.path[this.N*2] += add
	}
	if row+col == this.N-1 {
		this.path[this.N*2+1] += add
	}
	n := add * this.N
	if this.path[row] == n || this.path[this.N+col] == n ||
		this.path[this.N*2] == n || this.path[this.N*2+1] == n {
		return player
	}
	return 0
}

/**
 * Your TicTacToe object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Move(row,col,player);
 */
