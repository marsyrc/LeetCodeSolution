package greedy

import "sort"

//求最多有几个会议重叠的情况
func minMeetingRooms(intervals [][]int) int {
	nums := []int{}
	for _, v := range intervals {
		nums = append(nums, v[0]*10+2)
		nums = append(nums, v[1]*10+1)
	}
	sort.Ints(nums)
	res := 0
	need := 0
	for _, v := range nums {
		if v%10 == 2 {
			need++
		} else {
			need--
		}
		if need > res {
			res = need
		}
	}
	return res
}
