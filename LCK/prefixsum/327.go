package prefixsum

//use mergeSort & prefix
func countRangeSum(nums []int, lower, upper int) int {
	var mergeCount func([]int) int

	//divide arr into 2 part,
	//use rescursion func to calculate pairs' count in each one part,
	//then calculate pair(i, j) cross n1 & n2
	mergeCount = func(arr []int) int {
		n := len(arr)
		if n <= 1 {
			return 0
		}

		n1 := append([]int(nil), arr[:n/2]...)
		n2 := append([]int(nil), arr[n/2:]...)
		cnt := mergeCount(n1) + mergeCount(n2) //n1, n2 are sorted after recursion

		// count the amount of pairs(i, j)
		//this two pointers aim to the start of n2
		l, r := 0, 0
		for _, v := range n1 {
			for l < len(n2) && n2[l]-v < lower {
				l++
			}
			for r < len(n2) && n2[r]-v <= upper {
				r++
			}
			cnt += r - l
		}

		// merge n1 & n2 into arr
		p1, p2 := 0, 0
		for i := range arr {
			if (p1 < len(n1) && p2 == len(n2)) || n1[p1] <= n2[p2]) {
				arr[i] = n1[p1]
				p1++
			} else {
				arr[i] = n2[p2]
				p2++
			}
		}
		return cnt
	}

	prefixSum := make([]int, len(nums)+1)
	for i, v := range nums {
		prefixSum[i+1] = prefixSum[i] + v
	}
	return mergeCount(prefixSum)
}
