package prefixsum

import "math"

// 一开始以为是石子游戏III的区间dp方法，
// 后面发现要固定拿k个数字，所以剩下的数字连续且固定，
// 所以用总数减去最小化的连续窗口就可以。
// 思路类似于环形数组最大子序和（继承于经典的最大子序和问题），有兴趣可以去做一下呢
// 下面提供的算法是前缀和，当k和n接近的时候，滑动窗口的方法更快并且常数空间，
//但是前缀和更酷，我就要前缀和。
func maxScore(cardPoints []int, k int) int {
	n := len(cardPoints)
	winSize := n - k
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + cardPoints[i-1]
	}
	sum := math.MaxInt32
	res := 0
	for i := 0; i <= k; i++ {
		sum = min(sum, prefix[i+winSize]-prefix[i])
	}
	res = prefix[n] - sum
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
