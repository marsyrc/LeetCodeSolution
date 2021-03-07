package bag

import "math"

func closestCost(baseCosts []int, toppingCosts []int, target int) int {
	can := make([]bool, 20001)
	for _, b := range baseCosts {
		can[b] = true
	}
	top := []int{}
	top = append(top, toppingCosts...)
	top = append(top, toppingCosts...)
	//01背包
	for _, t := range top {
		for i := 20000; i >= t; i-- {
			can[i] = can[i] || can[i-t]
		}
	}
	minGap := math.MaxInt32
	ans := 0
	for i := 1; i <= 20000; i++ {
		if can[i] && abs(i-target) < minGap {
			ans = i
			minGap = abs(i - target)
		}
	}
	return ans
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
