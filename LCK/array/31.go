package array

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	j := i + 1
	if i >= 0 {
		k := n - 1
		for ; k >= j; k-- {
			if nums[k] > nums[i] {
				break
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	reverse(nums, j, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
