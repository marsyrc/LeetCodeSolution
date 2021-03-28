package resevoirsampling

import (
	"math/rand"
	"time"
)

//常数空间，线性时间
type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	return Solution{
		nums: nums,
	}
}

func (this *Solution) Pick(target int) int {
	rand.Seed(time.Now().UnixNano())
	count := 1
	res := -1
	for i, v := range this.nums {
		if v == target {
			if rand.Intn(count) == 0 {
				res = i
			}
			count++
		}
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Pick(target);
 */
