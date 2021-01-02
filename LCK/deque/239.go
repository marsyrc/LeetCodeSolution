package deque

//Monotonic queue
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < 2 {
		return nums
	}
	res := []int{}
	q := []int{}
	for i := 0; i < len(nums); i++ {
		//let q[0] be the index of max element in window
		for len(q) > 0 && nums[q[len(q)-1]] <= nums[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)

		if q[0] <= i-k {
			q = q[1:]
		}
		if i+1 >= k {
			res = append(res, nums[q[0]])
		}
	}
	return res
}
