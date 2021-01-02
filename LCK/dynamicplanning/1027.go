package dynamicplanning

import "math"

//dp considering two index, similiar from 873.go
//but here A is not increasing, demanding us to update indexMap of j
func longestArithSeqLength(A []int) int {
	if len(A) < 3 {
		return 0
	}

	dp := make([][]int, len(A))
	for i := range dp {
		dp[i] = make([]int, len(A))
	}
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			dp[i][j] = 2
		}
	}

	index := make(map[int]int)
	for i, v := range A {
		index[v] = i
	}

	res := 0
	for k := 0; k < len(A); k++ {
		for j := 0; j < k; j++ {
			if i, ok := index[2*A[j]-A[k]]; ok && i < j {
				dp[j][k] = dp[i][j] + 1
			}
			res = max(res, dp[j][k])
			//WATCH OUT : update the index of A[j], make it nearest from A[k]
			index[A[j]] = j
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
