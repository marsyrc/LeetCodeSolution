package mathproblem

import (
	"math/bits"
)

func pathInZigZagTree(label int) []int {
	res := []int{}

	for label != 1 {
		res = append(res, label)
		label >>= 1
		l := bits.Len(uint(label))
		label ^= (1 << (l - 1)) - 1
	}
	res = append(res, 1)
	reverse(res)
	return res
}

func reverse(nums []int) {
	n := len(nums)
	l, r := 0, n-1
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

//14 : 1110

// 1
// 11
// 100
// 1110
