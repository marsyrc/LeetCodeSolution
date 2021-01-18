package array

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	dir := 1
	row, col := 0, 0
	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0
	var res []int
	for top <= bottom && left <= right {
		res = append(res, matrix[row][col])
		switch dir {
		case 1:
			if col == right {
				dir = 2
				top++
				row++
				continue
			}
			col++
		case 2:
			if row == bottom {
				dir = 3
				right--
				col--
				continue
			}
			row++
		case 3:
			if col == left {
				dir = 4
				bottom--
				row--
				continue
			}
			col--
		case 4:
			if row == top {
				dir = 1
				left++
				col++
				continue
			}
			row--
		}
	}
	return res
}
