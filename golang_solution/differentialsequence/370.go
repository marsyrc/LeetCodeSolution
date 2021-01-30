package differentialsequence

func getModifiedArray(length int, updates [][]int) []int {
	diff := make([]int, length)
	for _, v := range updates {
		diff[v[0]] += v[2]
		if v[1]+1 < length {
			diff[v[1]+1] -= v[2]
		}
	}
	res := make([]int, length)
	res[0] = diff[0]
	for i := 1; i < length; i++ {
		res[i] = res[i-1] + diff[i]
	}
	return res
}
