package dynamicplanning

import "math"

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	for i := len(costs) - 2; i >= 0; i-- {
		costs[i][0] += min(costs[i+1][1], costs[i+1][2])
		costs[i][1] += min(costs[i+1][0], costs[i+1][2])
		costs[i][2] += min(costs[i+1][1], costs[i+1][0])
	}
	return min(costs[0][0], costs[0][1], costs[0][2])
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
