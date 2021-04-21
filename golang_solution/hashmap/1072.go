package hashmap

import "strconv"

func maxEqualRowsAfterFlips(matrix [][]int) int {
	hm := make(map[string]int)
	for i := range matrix {
		temp := ""
		if matrix[i][0] == 0 {
			for j := range matrix[i] {
				temp += strconv.Itoa(matrix[i][j])
			}
		} else {
			for j := range matrix[i] {
				temp += strconv.Itoa(matrix[i][j] ^ 1)
			}
		}
		hm[temp]++
	}
	res := -1
	for _, cnt := range hm {
		if cnt > res {
			res = cnt
		}
	}
	return res
}
