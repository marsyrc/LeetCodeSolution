package mathproblem

func minimumBoxes(n int) int {
	//towards perfect situation
	// 0 : 0
	// 1 : 0 + 1
	// 2 : 1 + 3
	// 3 : 4 + 6
	// 4 : 10 + 10
	// n * (n + 1) / 2, n = depth
	before, depth, total := 0, 0, 0
	for total < n {
		depth++
		total += depth * (depth + 1) / 2
		before += depth
	}
	//回退到上一个完美状态
	before -= depth
	total -= depth * (depth + 1) / 2
	ans := before
	//多出的方块总数/地面层加的 = 多出的数量 / n * (n + 1) / 2
	//以当前n = 4, ans = 3完美状态距离
	//n = 5 : ans = 3 + 1
	//n = 6 : ans = 3 + 2
	//n = 7 : ans = 3 + 2
	//n = 8 : ans = 3 + 3
	//n = 9 : ans = 3 + 3
	//n = 10 : ans = 3 + 3
	for i := 1; i <= depth; i++ {
		if total >= n {
			break
		} else {
			total += i
		}
		ans++
	}
	return ans
}
