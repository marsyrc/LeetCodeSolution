package array

func firstMissingPositive(nums []int) int {
	n := len(nums)

	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}

	//put nums[i] into somewhere it belongs
	//nums[nums[i] - 1] = nums[i]
	for i := 0; i < n; i++ {
		for nums[i]-1 >= 0 && nums[i]-1 < n && nums[i] != nums[nums[i]-1] {
			swap(i, nums[i]-1)
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}
