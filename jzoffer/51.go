package jianzhioffer

func reversePairs(nums []int) int {
	return mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, left, right int) int {
	if left >= right {
		return 0
	}
	mid := (left + right) / 2
	cnt := 0
	cnt += mergeSort(nums, left, mid) + mergeSort(nums, mid+1, right)
	i, j := left, mid+1
	tmp := []int{}
	for i <= mid && j <= right {
		if nums[i] <= nums[j] {
			tmp = append(tmp, nums[i])
			i++
			cnt += j - mid - 1
		} else {
			tmp = append(tmp, nums[j])
			j++
		}
	}

	for ; i <= mid; i++ {
		tmp = append(tmp, nums[i])
		cnt += right - mid
	}
	for ; j <= right; j++ {
		tmp = append(tmp, nums[j])
	}

	for i := left; i <= right; i++ {
		nums[i] = tmp[i-left]
	}
	return cnt
}
