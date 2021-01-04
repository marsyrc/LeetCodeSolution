package prefixsum

func subarraySum(nums []int, k int) int {
	cnt := 0
	cur := 0
	dict := make(map[int]int)
	dict[cur]++

	for _, v := range nums {
		//calculate current prefixSum
		cur += v
		//search in the map whether (prefix - k) in dict yet
		if u, ok := dict[cur-k]; ok {
			cnt += u
		}
		//insert current prefixSum into dict
		dict[cur]++
	}
	return cnt
}
