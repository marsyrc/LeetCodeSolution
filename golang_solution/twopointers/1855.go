package twopointers

func maxDistance(nums1 []int, nums2 []int) int {
	res := 0
	n1, n2 := len(nums1), len(nums2)
	i := 0
	for j := 0; j < n2; j++ {
		for i < n1 && nums1[i] > nums2[j] {
			i++
		}
		if i < n1 {
			res = max(res, j-i)
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//二分法
/*
func maxDistance(nums1 []int, nums2 []int) int {
    res := 0
    for i := 0; i < len(nums1); i++ {
        res = max(res, bs(nums2, nums1[i]) - i)
    }
    return res
}

func bs(nums []int, t int) int {
    l, r := 0, len(nums) - 1
    for l < r {
        mid := (l + r) / 2
        if nums[mid] >= t {
            l = mid + 1
        } else {
            r = mid
        }
    }
    if nums[l] < t {
        return l - 1
    }
    return l
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
*/
