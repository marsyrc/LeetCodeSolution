package array

func wiggleSort(nums []int) {
	less := true
	for i := 0; i < len(nums)-1; i++ {
		if less {
			if nums[i] > nums[i+1] {
				swap(nums, i, i+1)
			}
		} else {
			if nums[i] < nums[i+1] {
				swap(nums, i, i+1)
			}
		}
		less = !less
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
