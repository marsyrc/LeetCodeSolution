package array

func sortColors(nums []int) {
	//netherland banner alg
	l, r := 0, len(nums)-1
	cur := 0
	for cur <= r {
		if nums[cur] == 0 {
			swap(nums, cur, l)
			l++
		} else if nums[cur] == 2 {
			swap(nums, cur, r)
			r--
			cur--
		}
		cur++
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
