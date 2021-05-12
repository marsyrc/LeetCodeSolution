package array

func getMinSwaps(ns string, k int) int {
	nums := make([]int, len(ns))
	origin := make([]int, len(ns))
	for i := 0; i < len(ns); i++ {
		nums[i] = int(ns[i] - '0')
	}
	copy(origin, nums)
	for i := 0; i < k; i++ {
		nextPermutation(nums)
	}
	cnt := 0
	for i := 0; i < len(ns); i++ {
		if origin[i] != nums[i] {
			j := i + 1
			for origin[j] != nums[i] {
				j++
			}
			for j != i {
				swap(origin, j-1, j)
				cnt++
				j--
			}
		}
	}
	return cnt
}

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	j := i + 1
	if i >= 0 {
		k := n - 1
		for ; k >= j; k-- {
			if nums[k] > nums[i] {
				break
			}
		}
		nums[i], nums[k] = nums[k], nums[i]
	}
	reverse(nums, j, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}
