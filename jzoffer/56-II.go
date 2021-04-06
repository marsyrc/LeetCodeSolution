package jianzhioffer

func singleNumber(nums []int) int {
	var res int
	for i := 0; i < 32; i++ {
		tmp := 0
		for j := 0; j < len(nums); j++ {
			tmp += (nums[j] >> i) & 1
		}
		res |= (tmp % 3) << i
	}
	return res
}
