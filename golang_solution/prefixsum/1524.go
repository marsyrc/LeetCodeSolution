package prefixsum

func numOfSubarrays(arr []int) int {
	const mod int = int(1e9 + 7)
	even, odd, res, cur := 1, 0, 0, 0
	for i := 0; i < len(arr); i++ {
		cur += arr[i]
		if cur&1 == 1 {
			res += even
			odd++
		} else {
			res += odd
			even++
		}
		res %= mod
	}
	return res
}
