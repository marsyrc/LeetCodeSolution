package design

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	origin []int
	r      *rand.Rand
}

func Constructor(nums []int) Solution {
	return Solution{
		origin: nums,
		r:      rand.New(rand.NewSource(time.Now().Unix())),
	}
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.origin
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	n := len(this.origin)
	res := make([]int, n)
	copy(res, this.origin)
	fmt.Println(res)

	for i := 0; i < n; i++ {
		idx := n - i
		nextRange := this.r.Intn(idx)
		b := i + nextRange
		res[i], res[b] = res[b], res[i]
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
