package jianzhioffer

func search(nums []int, target int) int {
	l, r := searchLeftBoundary(nums, target), searchRightBoundary(nums, target)
	if l == -1 && r == -1 {
		return 0
	}
	return r - l + 1
}

func searchLeftBoundary(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			r = mid - 1
		}
	}
	if l >= len(nums) || nums[l] != target {
		return -1
	}
	return l
}

func searchRightBoundary(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if target > nums[mid] {
			l = mid + 1
		} else if target < nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if r < 0 || nums[r] != target {
		return -1
	}
	return r
}

//53-II.go
//一个长度为n-1的递增排序数组中的所有数字都是唯一的，并且每个数字都在范围0～n-1之内。
//在范围0～n-1内的n个数字中有且只有一个数字不在该数组中，请找出这个数字。

func missingNumber(nums []int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] != mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}
