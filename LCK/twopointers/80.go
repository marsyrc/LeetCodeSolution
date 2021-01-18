package twopointers

func removeDuplicates(nums []int) int {
	maxRepeat := 2
	slow := maxRepeat - 1
	for fast := maxRepeat; fast < len(nums); fast++ {
		if nums[slow-maxRepeat+1] != nums[fast] {
			slow++
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
