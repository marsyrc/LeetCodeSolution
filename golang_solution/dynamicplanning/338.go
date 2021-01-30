package dynamicplanning

func countBits(num int) []int {
	ans := make([]int, num+1)
	i := 0
	b := 1
	for b <= num {
		for i < b && i+b <= num {
			ans[i+b] = 1 + ans[i]
			i++
		}
		i = 0
		b *= 2
	}
	return ans
}
