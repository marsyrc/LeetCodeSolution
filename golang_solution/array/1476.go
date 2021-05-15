package array

type SubrectangleQueries struct {
	rec    [][]int
	change [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{
		rec:    rectangle,
		change: [][]int{},
	}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.change = append(this.change, []int{row1, col1, row2, col2, newValue})
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	for i := len(this.change) - 1; i >= 0; i-- {
		if this.change[i][0] <= row && this.change[i][1] <= col &&
			this.change[i][2] >= row && this.change[i][3] >= col {
			return this.change[i][4]
		}
	}
	return this.rec[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
