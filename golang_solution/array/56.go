package array

import "sort"

//区间题经典
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0] || intervals[i][0] == intervals[j][0] && intervals[i][1] > intervals[j][1]
	})
	l, r := intervals[0][0], intervals[0][1]
	res := [][]int{{l, r}}
	for i := 1; i < len(intervals); i++ {
		for i < len(intervals) && intervals[i][0] == intervals[i-1][0] {
			i++
		}
		if i == len(intervals) {
			return res
		}
		if intervals[i][0] <= r {
			if intervals[i][1] <= r {
				continue
			} else {
				res = res[:len(res)-1]
				r = intervals[i][1]
				res = append(res, []int{l, r})
			}
		} else {
			l, r = intervals[i][0], intervals[i][1]
			res = append(res, []int{l, r})
		}
	}
	return res
}
