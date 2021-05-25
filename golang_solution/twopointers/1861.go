package twopointers

func rotateTheBox(box [][]byte) [][]byte {
	m, n := len(box), len(box[0])
	for i := 0; i < m; i++ {
		last := n - 1
		for j := n - 1; j >= 0; j-- {
			if box[i][j] == '*' {
				last = j - 1
			} else if box[i][j] == '#' {
				if last > j {
					box[i][last] = '#'
					box[i][j] = '.'
					last--
				} else {
					last = j - 1
				}
			}
		}
	}

	res := make([][]byte, n)
	for i := range res {
		res[i] = make([]byte, m)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[j][m-i-1] = box[i][j]
		}
	}
	return res
}
