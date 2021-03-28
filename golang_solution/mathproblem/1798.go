package mathproblem

import "sort"

func getMaximumConsecutive(coins []int) int {
	sort.Ints(coins)
	x := 0
	for _, coin := range coins {
		if coin > x+1 {
			break
		}
		x += coin
	}
	return x + 1
}
