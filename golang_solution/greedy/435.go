package greedy

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	if len(intervals) < 2 {
		return 0
	}
	pivot := 0
	res := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[pivot][1] > intervals[i][0] {
			if intervals[pivot][1] > intervals[i][1] {
				pivot = i
			}
			res++
		} else {
			pivot = i
		}
	}
	return res
}
