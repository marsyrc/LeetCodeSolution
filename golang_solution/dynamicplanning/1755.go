package dynamicplanning

import (
	"math"
	"sort"
)

//状态压缩dp,将n个数在不在子数组的状态压缩在大小为2 ^ n的数组中
//如果当前i满足 i&(1<<j) != 0,说明当前sum有nums[j]
//即Sum[i] = Sum[i-1<<j] + nums[j]
//分成两半处理是满足时间限制
func minAbsDifference(nums []int, goal int) int {
	n := len(nums)
	half := n / 2
	ls, rs := half, n-half
	lSum, rSum := make([]int, 1<<ls), make([]int, 1<<rs)
	for i := 1; i < (1 << ls); i++ {
		for j := 0; j < ls; j++ {
			if i&(1<<j) != 0 {
				lSum[i] = lSum[i-1<<j] + nums[j]
				break
			}
		}
	}
	for i := 1; i < (1 << rs); i++ {
		for j := 0; j < rs; j++ {
			if i&(1<<j) != 0 {
				rSum[i] = rSum[i-1<<j] + nums[j+ls]
				break
			}
		}
	}
	sort.Ints(rSum)
	sort.Ints(lSum)
	res := math.MaxInt32
	for _, n := range rSum {
		res = min(res, abs(goal-n))
	}
	for _, n := range lSum {
		res = min(res, abs(goal-n))
	}
	i, j := 0, len(rSum)-1
	for i < len(lSum) && j >= 0 {
		s := lSum[i] + rSum[j]
		res = min(res, abs(goal-s))
		if s > goal {
			j--
		} else {
			i++
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
