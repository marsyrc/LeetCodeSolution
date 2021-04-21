package recursion

import "sort"

func getKth(lo int, hi int, k int) int {
	hash := make([]int, 1001)
	hash[1] = 1

	arr := make([]int, hi-lo+1)
	for i := lo; i <= hi; i++ {
		arr[i-lo] = i
		getQ(i, hash)
	}

	sort.Slice(arr, func(i, j int) bool {
		if hash[arr[i]] == hash[arr[j]] {
			return arr[i] < arr[j]
		}
		return hash[arr[i]] < hash[arr[j]]
	})
	return arr[k-1]
}

func getQ(num int, hash []int) int {
	if num <= 1000 && hash[num] > 0 {
		return hash[num]
	}
	r := 0
	if num%2 == 0 {
		r = getQ(num/2, hash) + 1
	} else {
		r = getQ(3*num+1, hash) + 1
	}
	if num <= 1000 {
		hash[num] = r
	}
	return r
}
