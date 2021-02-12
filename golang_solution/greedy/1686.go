package greedy

import "sort"

func stoneGameVI(aliceValues []int, bobValues []int) int {
	pairs := []pair{}
	for i := range aliceValues {
		pairs = append(pairs, pair{aliceValues[i] + bobValues[i], i})
	}
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].dis > pairs[j].dis
	})
	sum1, sum2 := 0, 0
	for i, p := range pairs {
		if i%2 == 0 {
			sum1 += aliceValues[p.i]
		} else {
			sum2 += bobValues[p.i]
		}
	}
	if sum1 > sum2 {
		return 1
	} else if sum2 > sum1 {
		return -1
	}
	return 0
}

type pair struct {
	dis int
	i   int
}
