package greedy

func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	for i := range nums1 {
		sum1 += nums1[i]
	}
	for i := range nums2 {
		sum2 += nums2[i]
	}
	if sum1 > sum2 {
		return minOperations(nums2, nums1)
	}
	diff := sum2 - sum1
	freq := make([]int, 6)
	for _, num := range nums1 {
		freq[6-num]++
	}
	for _, num := range nums2 {
		freq[num-1]++
	}
	res := 0
	for i := 5; i >= 1 && diff > 0; i-- {
		for freq[i] > 0 && diff > 0 {
			res++
			freq[i]--
			diff -= i
		}
	}
	if diff > 0 {
		return -1
	}
	return res
}
