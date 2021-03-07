package stack

func nextGreaterElements(nums []int) []int {
	res := make([]int, len(nums))
	var s []int
	n := len(nums)
	if n == 1 {
		return []int{-1}
	}
	for i := 2*len(nums) - 1; i >= 0; i-- {
		for len(s) > 0 && s[len(s)-1] <= nums[i%n] {
			s = s[:len(s)-1]
		}
		if len(s) == 0 {
			res[i%n] = -1
		} else {
			res[i%n] = s[len(s)-1]
		}
		s = append(s, nums[i%n])
	}
	return res
}
