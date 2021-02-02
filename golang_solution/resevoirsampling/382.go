package resevoirsampling

import (
	"math/rand"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

type Solution struct {
	head *ListNode
	r    *rand.Rand
}

/** @param head The linked list's head.
  Note that the head is guaranteed to be not null, so it contains at least one node. */
func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
		r:    rand.New(rand.NewSource(time.Now().Unix())),
	}
}

/** Returns a random node's value. */
func (this *Solution) GetRandom() int {
	res := 0
	count := 1
	cur := this.head
	for cur != nil {
		randN := this.r.Intn(count)
		if 0 == randN {
			res = cur.Val
		}
		cur = cur.Next
		count++
	}
	return res
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
