package mathproblem

func addNegabinary(arr1 []int, arr2 []int) []int {
	ans := make([]int, 0)
	w := 0
	n := max(len(arr1), len(arr2))
	for i := 0; i < n; i++ {
		n1, n2 := 0, 0
		if len(arr1) > i {
			n1 = arr1[len(arr1)-i-1]
		}
		if len(arr2) > i {
			n2 = arr2[len(arr2)-i-1]
		}
		if n1+n2-w < 2 && n1+n2-w >= 0 {
			ans = append([]int{n1 + n2 - w}, ans...)
			w = 0
		} else if n1+n2-w == 2 {
			w = 1
			ans = append([]int{0}, ans...)
		} else if n1+n2-w == -1 {
			w = -1
			ans = append([]int{1}, ans...)
		} else if n1+n2-w == 3 {
			w = 1
			ans = append([]int{1}, ans...)
		}
	}
	if w == 1 {
		ans = append([]int{1, 1}, ans...)
	}
	i := 0
	for i < len(ans) && ans[i] == 0 {
		i++
	}
	if i == len(ans) {
		return []int{0}
	}
	return ans[i:]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
