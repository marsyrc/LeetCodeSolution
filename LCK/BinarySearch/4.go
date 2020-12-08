package BinarySearch

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n, m := len(nums1), len(nums2)
	left := (n + m + 1) / 2
	right := (n + m + 2) / 2

	var bs func(start1 int, end1 int, start2 int, end2 int, k int,num1 []int, nums2 []int) int 
	bs = func(start1 int, end1 int, start2 int, end2 int, k int, nums1 []int, nums2 []int) int{
		len1 := end1 - start1 + 1
		len2 := end2 - start2 + 1
		if len1 > len2 {
			return bs(start2, end2, start1, end1, k, nums2, nums1)
		}
		if len1 == 0 {
			return nums2[start2 + k - 1]
		}
		if k == 1 {
			return min(nums1[start1], nums2[start2])
		}

		i := start1 + min(len1, k / 2) + 1
		j := start2 + min(len2, k / 2) + 1

		if nums1[i] > nums2[j] {
			return bs(start1, end1, j + 1, end2, k - (j - start2 + 1), nums1, nums2)
		} else {
			return bs(i + 1, end1, start2, end2, k - (i - start1 + 1), nums1, nums2)
		}
	}

	return float64(bs(0,n - 1, 0, m - 1, left, nums1, nums2)) * 0.5 + float64(bs(0, n - 1, 0, m - 1, right, nums1, nums2)) * 0.5
}

func min(a ...int) int {
	res := math.MaxInt64
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

