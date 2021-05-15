package prefixsum

func countTriplets(arr []int) int {
	n := len(arr)
	p := make([]int, n)
	p[0] = arr[0]
	for i := 1; i < n; i++ {
		p[i] = p[i-1] ^ arr[i]
	}

	res := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j; k < n; k++ {
				if i == 0 {
					if p[k]^p[j-1] == p[j-1] {
						res++
					}
					continue
				}
				if p[k]^p[j-1] == p[j-1]^p[i-1] {
					res++
				}
			}
		}
	}
	return res
}
