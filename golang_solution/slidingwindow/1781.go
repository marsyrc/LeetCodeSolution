package slidingwindow

import "math"

func beautySum(s string) int {
	res := 0
	for i := 0; i < len(s); i++ {
		cnt := make([]int, 26)
		for j := i; j < len(s); j++ {
			curC := int(s[j] - 'a')
			minCnt := math.MaxInt32
			maxCnt := 0
			cnt[curC]++
			for _, c := range cnt {
				if c != 0 {
					maxCnt = max(maxCnt, c)
					minCnt = min(minCnt, c)
				}
			}
			res += maxCnt - minCnt
		}
	}
	return res
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
