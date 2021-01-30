package greedy

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var ans []int
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		var numsA []int
		var numsB []int
		numsA = getMaxSonArray(nums1, i)
		numsB = getMaxSonArray(nums2, k-i)
		if less(ans, merge(numsA, numsB)) {
			ans = merge(numsA, numsB)
		}
	}
	return ans
}

func getMaxSonArray(nums []int, k int) []int {
	s := make([]int, k)
	n := len(nums)
	for i := 0; i < n; i++ {
		for len(s) != 0 && s[len(s)-1] < nums[i] && len(nums)-i+len(s) > k {
			s = s[:len(s)-1]
		}
		if len(s) < k {
			s = append(s, nums[i])
		}
	}
	return s
}

//使用less逻辑来判断归并方向，传统原地归并有问题
func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		if less(a, b) {
			merged[i], b = b[0], b[1:]
		} else {
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}

func less(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}
