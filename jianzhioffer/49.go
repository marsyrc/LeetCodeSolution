package jianzhioffer

import (
	"container/heap"
	"math"
)

func nthUglyNumber(n int) int {
	uglyNum := -1
	var h IntHeap
	heap.Init(&h)
	heap.Push(&h, 1)
	for n > 0 {
		for h[0] == uglyNum {
			heap.Pop(&h)
		}
		uglyNum = heap.Pop(&h).(int)

		if 2*uglyNum <= math.MaxInt32 {
			heap.Push(&h, 2*uglyNum)
		}
		if 3*uglyNum <= math.MaxInt32 {
			heap.Push(&h, 3*uglyNum)
		}
		if 5*uglyNum <= math.MaxInt32 {
			heap.Push(&h, 5*uglyNum)
		}
		n--
	}
	return uglyNum
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
