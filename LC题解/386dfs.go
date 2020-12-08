// 386. 字典序排数

// 给定一个整数 n, 返回从 1 到 n 的字典顺序。

// 例如，

// 给定 n =1 3，返回 [1,10,11,12,13,2,3,4,5,6,7,8,9] 。

// 请尽可能的优化算法的时间复杂度和空间复杂度。 输入的数据 n 小于等于 5,000,000。

func lexicalOrder(n int) []int {
	var res []int
	var dfs func(num int)
	dfs = func(num int) {
		if num > n {
			return
		}
		res = append(res, num)
		for i := 0; i <= 9; i++ {
			dfs(num*10 + i)
		}
	}
	for i := 1; i <= 9; i++ {
		dfs(i)
	}
	return res
}