package jianzhioffer

func majorityElement(nums []int) int {
	cand := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if cand == nums[i] {
			cnt++
		} else {
			cnt--
			if cnt == 0 {
				cnt = 1
				cand = nums[i]
			}
		}
	}
	return cand
}
