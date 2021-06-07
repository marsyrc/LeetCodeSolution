package differentialsequence

import "sort"

func maxSumRangeQuery(nums []int, requests [][]int) int {
	n := len(nums)
	diff := make([]int, n+1)
	for _, request := range requests {
		diff[request[0]]++
		diff[request[1]+1]--
	}
	for i := 1; i <= n; i++ {
		diff[i] += diff[i-1]
	}

	sort.Ints(nums)
	sort.Ints(diff)
	ans := 0
	p := int(1e9 + 7)
	for i := n; i >= 1 && diff[i] > 0; i-- {
		ans += diff[i] * nums[i-1]
		ans %= p
	}
	return ans
}
