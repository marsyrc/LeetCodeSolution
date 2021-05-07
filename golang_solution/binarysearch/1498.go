package binarysearch

import "sort"

const mod = 1e9 + 7

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	count := 0

	bs := func(t int) int {
		l, r := 0, len(nums)-1
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] > t {
				r = mid - 1
			} else {
				if mid == len(nums)-1 || nums[mid+1] > t {
					return mid
				} else {
					l = mid + 1
				}
			}
		}
		return -1
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > target/2+1 {
			break
		}
		j := bs(target - nums[i])
		if j != -1 && j >= i {
			count += int(myPow(2, j-i))
			count %= mod
		}
	}

	return count
}

func myPow(x int, n int) int {
	if n < 0 {
		return 1 / myPow(x, -n)
	}
	if n == 0 {
		return 1
	}
	y := myPow(x, n/2)
	if n%2 == 1 {
		return y * y * x % mod
	}
	return y * y % mod
}
