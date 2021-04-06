package jianzhioffer

func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	swap := func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for l < r {
		for l < len(nums)-1 && nums[l]%2 == 1 {
			l++
		}
		for r > 0 && nums[r]%2 == 0 {
			r--
		}
		if l < r {
			swap(l, r)
		}
	}
	return nums
}
