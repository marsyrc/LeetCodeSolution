package prefixsum

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	prefix := make([]int, 1+n)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
	}

	for i := 0; i < n; i++ {
		if i > 0 {
			res[i] += nums[i]*i - prefix[i]
			// fmt.Println(i, res[i])
		}
		if i < n-1 {
			res[i] += prefix[n] - prefix[i+1] - (n-i-1)*nums[i]
			// fmt.Println(i, res[i])
		}
	}
	return res
}
