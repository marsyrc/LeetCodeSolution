package slidingwindow

func subarraysWithKDistinct(A []int, K int) int {
	return helper(A, K) - helper(A, K-1)
}

func helper(A []int, k int) int {
	cnt := 0
	res := 0
	left, right := 0, 0
	freq := make([]int, len(A)+1)
	for right < len(A) {
		if freq[A[right]] == 0 {
			cnt++
		}
		freq[A[right]]++
		right++
		for cnt > k {
			freq[A[left]]--
			if freq[A[left]] == 0 {
				cnt--
			}
			left++
		}
		res += right - left
	}
	return res
}
