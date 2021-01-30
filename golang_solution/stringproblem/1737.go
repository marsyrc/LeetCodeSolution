package stringproblem

import "math"

//计数 + 前缀和
func minCharacters(a string, b string) int {
	cntA, cntB := [26]int{}, [26]int{}
	for _, c := range a {
		cntA[c-'a']++
	}
	for _, c := range b {
		cntB[c-'a']++
	}
	res := math.MaxInt32
	aSum, bSum := 0, 0
	aN, bN := len(a), len(b)
	for i := 0; i < 25; i++ {
		aSum += cntA[i]
		bSum += cntB[i]
		res = min(res, aN+bN-cntA[i]-cntB[i], aN-aSum+bSum, bN-bSum+aSum)
	}
	res = min(res, aN+bN-cntA[25]-cntB[25])
	return res
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if res > v {
			res = v
		}
	}
	return res
}
