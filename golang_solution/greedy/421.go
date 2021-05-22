package greedy

func findMaximumXOR(nums []int) int {
	res := 0
	mask := 0
	for i := 31; i >= 0; i-- {
		mask |= (1 << i)
		set := make(map[int]struct{})
		for _, num := range nums {
			set[num&mask] = struct{}{}
		}

		expected := (1 << i) | res
		for v := range set {
			if _, ok := set[expected^v]; ok {
				res = expected
				break
			}
		}
	}
	return res
}
