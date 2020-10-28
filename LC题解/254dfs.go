/**
254. 因子的组合

整数可以被看作是其因子的乘积。

例如：

8 = 2 x 2 x 2;
  = 2 x 4.

请实现一个函数，该函数接收一个整数 n 并返回该整数所有的因子组合。

注意：

    你可以假定 n 为永远为正数。
    因子必须大于 1 并且小于 n。

示例 1：

输入: 1
输出: []

示例 2：

输入: 37
输出: []

示例 3：

输入: 12
输出:
[
  [2, 6],
  [2, 2, 3],
  [3, 4]
]
*/

func getFactors(n int) [][]int {
	path := []int{}
	var ans [][]int

	var dfs func(n int)
	dfs = func(n int) {
		if n <= 1 {
			return
		}

		for i := 2; i*i <= n; i++ {
			if len(path) != 0 {
				if i < path[len(path)-1] || n/i < path[len(path)-1] {
					continue
				}
			}

			if n%i == 0 {
				path = append(path, i, n/i)
				tmp := make([]int, len(path))
				copy(tmp, path)
				ans = append(ans, tmp)
				path = path[:len(path)-1]
				dfs(n / i)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(n)
	return ans
}