package array

import "math"

func insert(intervals [][]int, newInterval []int) (res [][]int) {
	left, right := newInterval[0], newInterval[1]
	n := len(intervals)
	i := 0
	//skip intervals on left of new
	for i < n && intervals[i][1] < left {
		res = append(res, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= right {
		left = min(intervals[i][0], left)
		right = max(intervals[i][1], right)
		i++
	}
	res = append(res, []int{left, right})

	for i < n {
		res = append(res, intervals[i])
		i++
	}
	return
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
