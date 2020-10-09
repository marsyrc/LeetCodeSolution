/**
剑指 Offer 13. 机器人的运动范围

地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，
它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。
例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，
因为3+5+3+8=19。请问该机器人能够到达多少个格子？

示例 1：

输入：m = 2, n = 3, k = 1
输出：3

示例 2：

输入：m = 3, n = 1, k = 0
输出：1


*/
type pair struct {
	x, y int
}

var directions = []pair{{1, 0}, {0, 1}}

func movingCount(m int, n int, k int) int {

	digitSum := func(x int) (res int) {
		for ; x > 0; x = x / 10 {
			res += x % 10
		}
		return
	}

	var check func(i, j int) bool
	check = func(i, j int) bool {
		sum := digitSum(i) + digitSum(j)
		if sum <= k {
			return true
		}
		return false
	}

	var queue = []pair{{0, 0}}
	counts := 1

	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}

	for len(queue) > 0 {
		size := len(queue)
		for i := size; i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			for _, dir := range directions {
				x := node.x + dir.x
				y := node.y + dir.y
				if x < m && y < n && x >= 0 && y >= 0 && !visited[x][y] {
					if !check(x, y) {
						continue
					} else {
						queue = append(queue, pair{x, y})
						visited[x][y] = true
						counts++
					}
				} else {
					continue
				}
			}
		}
	}
	return counts
}
