package stack

func sumSubarrayMins(arr []int) int {
	//求出每个值作为最小值存在的区间的个数 * 该值 加到结果中
	const mod int = 1e9 + 7
	//让arr所有元素都能出栈，所以在后面加个巨小的数字
	arr = append(arr, 0)
	s := []int{}
	ans := 0
	for i := 0; i < len(arr); i++ {
		for len(s) > 0 && arr[s[len(s)-1]] >= arr[i] {
			idx := s[len(s)-1]
			s = s[:len(s)-1]
			pre := idx + 1
			if len(s) != 0 {
				pre = idx - s[len(s)-1]
			}
			post := i - idx
			ans += arr[idx] * post * pre
			ans %= mod
		}
		s = append(s, i)
	}
	return ans
}
