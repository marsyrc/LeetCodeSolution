package heapsort

import "container/heap"

//this problem has better solutions
//see deque/239.go
func maxSlidingWindow(nums []int, k int) []int {
	h := maxHeap{}
	for i := 0; i < k; i++ {
		heap.Push(&h, pair{nums[i], i})
	}
	start, end := 0, k-1
	res := []int{}
	res = append(res, h[0].val)
	for end < len(nums)-1 {
		start++
		end++
		for len(h) > 0 && h[0].idx < start {
			heap.Pop(&h)
		}
		heap.Push(&h, pair{nums[end], end})
		res = append(res, h[0].val)
	}
	return res
}

type pair struct {
	val int
	idx int
}

type maxHeap []pair

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i].val > h[j].val }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *maxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
