package greedy

import "math"

func minimumTeachings(n int, languages [][]int, friendships [][]int) int {
	hasCommon := func(i, j int) bool {
		for _, l1 := range languages[i-1] {
			for _, l2 := range languages[j-1] {
				if l1 == l2 {
					return true
				}
			}
		}
		return false
	}
	notConnected := map[int]int{}
	for _, f := range friendships {
		if !hasCommon(f[0], f[1]) {
			notConnected[f[0]]++
			notConnected[f[1]]++
		}
	}
	mostUsedLang := make([]int, n+1)
	for p := range notConnected {
		for _, l := range languages[p-1] {
			mostUsedLang[l]++
		}
	}
	notConnectedCnt := len(notConnected)
	most := max(mostUsedLang...)
	return notConnectedCnt - most
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
