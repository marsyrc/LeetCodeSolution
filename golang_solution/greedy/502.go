package greedy

import (
	"container/heap"
	"fmt"
)

func findMaximizedCapital(k int, W int, Profits []int, Capital []int) int {
	h1 := MaxHeap{}
	h2 := MinHeap{}
	res := W
	for i := range Capital {
		cur := pair{Profits[i], Capital[i]}
		heap.Push(&h2, cur)
	}
	for k > 0 {
		fmt.Println("res", res)
		for len(h2) != 0 && h2[0].c <= res {
			top := heap.Pop(&h2).(pair)
			fmt.Println("cap:", top.c)
			heap.Push(&h1, top)
		}
		if len(h1) == 0 {
			return res
		} else {
			mostPro := heap.Pop(&h1).(pair)
			res += mostPro.p
			k--
		}
	}
	return res
}

type pair struct {
	p int
	c int
}

//大根堆，按利润排序
type MaxHeap []pair

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return h[i].p > h[j].p
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return res
}
func (h *MaxHeap) Push(x interface{}) {
	*h = append((*h), x.(pair))
}

//按资本排序
type MinHeap []pair

func (h MinHeap) Len() int {
	return len(h)
}
func (h MinHeap) Less(i, j int) bool {
	return h[i].c < h[j].c
}
func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MinHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return res
}
func (h *MinHeap) Push(x interface{}) {
	*h = append((*h), x.(pair))
}
