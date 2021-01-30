package heapsort

import "container/heap"

type t struct {
	//date : when apples rot
	//num : nums of apples
	date int
	num  int
}

type minHeap []t

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].date < h[j].date }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(t)) }
func (h *minHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func eatenApples(apples []int, days []int) (ans int) {
	h := minHeap{}
	for i := 0; i < len(days) || len(h) != 0; i++ {
		//product
		if i < len(days) && apples[i] != 0 {
			newT := t{
				date: days[i] + i,
				num:  apples[i],
			}
			heap.Push(&h, newT)
		}

		//update ts which out of date
		for len(h) != 0 && h[0].date == i {
			heap.Pop(&h)
		}

		//eat one
		if len(h) != 0 && h[0].num > 0 {
			ans++
			h[0].num--
			if h[0].num == 0 {
				heap.Pop(&h)
			}
		}
	}
	return
}
