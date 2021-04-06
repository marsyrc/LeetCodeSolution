package jianzhioffer

import "container/heap"

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
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

type MedianFinder struct {
	maxh *maxHeap
	minh *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	res := MedianFinder{
		maxh: new(maxHeap),
		minh: new(minHeap),
	}

	return res
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxh, num)
	heap.Push(this.minh, (heap.Pop(this.maxh)).(int))

	if len(*this.maxh) < len(*this.minh) {
		heap.Push(this.maxh, (heap.Pop(this.minh)).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(*this.minh) == len(*this.maxh) {
		return (float64((*this.maxh)[0]) + float64((*this.minh)[0])) / float64(2)
	} else {
		return float64((*this.maxh)[0])
	}
}
