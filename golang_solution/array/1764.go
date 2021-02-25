package array

func canChoose(groups [][]int, nums []int) bool {
	m, n := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == groups[m][n] {
			n++
			if n == len(groups[m]) {
				m++
				if m == len(groups) {
					return true
				}
				n = 0
			}
		} else {
			i -= n
			n = 0
		}
	}
	return false
}
