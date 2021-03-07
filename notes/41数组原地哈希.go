// 41. 缺失的第一个正数

// 给你一个未排序的整数数组，请你找出其中没有出现的最小的正整数。

// 示例 1:

// 输入: [1,2,0]
// 输出: 3

// 示例 2:

// 输入: [3,4,-1,1]
// 输出: 2

// 示例 3:

// 输入: [7,8,9,11,12]
// 输出: 1

func firstMissingPositive(nums []int) int {
	n := len(nums)
	var swap func(i, j int)
	swap = func(i, j int) {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i := 0; i < n; i++ {
		for nums[i]-1 >= 0 && nums[i]-1 < n && nums[nums[i]-1] != nums[i] {
			swap(nums[i]-1, i)
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return n + 1
}