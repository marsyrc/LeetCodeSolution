// res[i][j]表示轮到这一个人（无论是A或者B）选时，他能得到与另一个人最大的价值差
// 无论是B想要缩短价值差，或者A想要扩大价值差，其本质就是这次取法能得到与对手最大的价值差
// 所以这次DP也用类似的状态。

// 又因为每次删掉之后的得分是区间累加和，所以又涉及到区间和，所以需要维护
// dp[i][j] : 表示从 i 到 j 的区间和
// res[i][j]:表示轮到这一个人选时，石头剩下 i 到 j ，他能获得的最大价值差

// 接下来推状态转移方程：
//     最初始的时候：i == j ，res[i][j] = 0，因为删了之后没得取
//     当 j - i == 1时，res[i][j] = max(stones[i], stones[j])，因为我要利益最大化，我肯定删掉一个较小的石头，取最大得分，反正下一个人是没分的
//     当j - i > 1时， res[i][j] = max(dp[i + 1][j] - res[i + 1][j], dp[i][j - 1] - res[i][j - 1]),
//     本次选取，我只能从左端删或者从右端删，
//     从左端删，在剩下的石头中（不考虑前面的），下一个人与我的得分差值是 res[i + 1][j]
//     从右端删，在剩下的石头中（不考虑前面的），下一个人与我的得分差值是 res[i][j - 1]
//     所以我考虑本次差值最大，那么本次删掉的时候我能得到的最大差值应该是删掉一个之后得到的区间和，减去对手下次选时与我的差值
//     所以转移方程为： res[i][j] = max(dp[i + 1][j] - res[i + 1][j], dp[i][j - 1] - res[i][j - 1])
func stoneGameVII(stones []int) int {
	n := len(stones)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if i == j {
				dp[i][j] = stones[i]
			} else {
				dp[i][j] = stones[j] + dp[i][j-1]
			}
		}
	}
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			if j-i == 1 {
				res[i][j] = max(stones[i], stones[j])
			} else {
				res[i][j] = max(dp[i+1][j]-res[i+1][j], dp[i][j-1]-res[i][j-1])
			}
		}
	}
	return res[0][n-1]
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}