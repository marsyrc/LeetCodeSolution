package heapsort

import (
	"container/heap"
)

func lastStoneWeight(stones []int) int {
	h := maxHeap{}
	for _, v := range stones {
		heap.Push(&h, v)
	}
	for len(h) > 1 {
		y, x := heap.Pop(&h).(int), heap.Pop(&h).(int)
		if y != x {
			heap.Push(&h, y-x)
		}
	}
	if len(h) != 0 {
		return h[0]
	}
	return 0
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
