package dynamicplanning

import "sort"

//this problem also has dp solutions, see dp/368.go
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	level := make([]int, n)
	father := make([]int, n)

	for i := 0; i < n; i++ {
		father[i] = i
	}

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && level[i] < level[j]+1 {
				level[i] = level[j] + 1
				father[i] = j
			}
		}
	}

	maxLevel := -1
	x := -1
	for i, v := range level {
		if v > maxLevel {
			x = i
			maxLevel = v
		}
	}
	res := []int{nums[x]}
	for x != father[x] {
		x = father[x]
		res = append(res, nums[x])
	}
	return res
}
