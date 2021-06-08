package greedy

import "sort"

func maxBoxesInWarehouse(boxes []int, warehouse []int) int {
	sort.Ints(boxes)
	n := len(warehouse)
	idx := n - 1
	for i := 1; i < len(warehouse); i++ {
		warehouse[i] = min(warehouse[i], warehouse[i-1])
	}
	res := 0
	b := 0
	for b < len(boxes) && idx >= 0 {
		cur := boxes[b]
		if cur > warehouse[idx] {
			idx--
		} else {
			idx--
			res++
			b++
		}
	}
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
