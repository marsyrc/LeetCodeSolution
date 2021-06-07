package differentialsequence

func minNumberOperations(target []int) int {
	res := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			res += target[i] - target[i-1]
		}
	}
	return res
}
