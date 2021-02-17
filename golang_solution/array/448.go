package array

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	for _, v := range nums {
		nums[(v-1)%n] += n
	}
	ans := []int{}
	for i := 0; i < n; i++ {
		if nums[i] <= n {
			ans = append(ans, i+1)
		}
	}
	return ans
}
