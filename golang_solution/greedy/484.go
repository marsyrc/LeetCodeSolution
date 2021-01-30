package greedy

func findPermutation(s string) []int {
	ans := make([]int, len(s)+1)
	for i := 1; i <= len(s)+1; i++ {
		ans[i-1] = i
	}
	i := 1
	for i <= len(s) {
		j := i
		for i <= len(s) && s[i-1] == 'D' {
			i++
		}
		reverse(ans, j-1, i-1)
		i++
	}
	return ans
}

func reverse(arr []int, i, j int) {
	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}
