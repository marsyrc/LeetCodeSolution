package heapsort

import (
	"container/heap"
	"fmt"
)

func getNumberOfBacklogOrders(orders [][]int) int {
	s, b := sellHeap{}, buyHeap{}
	for _, order := range orders {
		//buyOrder
		if order[2] == 0 {
			for order[1] > 0 && len(s) > 0 && s[0][0] <= order[0] {
				top := heap.Pop(&s).([]int)
				temp := min(top[1], order[1])
				top[1] -= temp
				order[1] -= temp
				if top[1] > 0 {
					heap.Push(&s, top)
				}
			}
			if order[1] > 0 {
				heap.Push(&b, order)
			}
		} else {
			for order[1] > 0 && len(b) > 0 && b[0][0] >= order[0] {
				top := heap.Pop(&b).([]int)
				temp := min(top[1], order[1])
				top[1] -= temp
				order[1] -= temp
				if top[1] > 0 {
					heap.Push(&b, top)
				}
			}
			if order[1] > 0 {
				heap.Push(&s, order)
			}
		}
		fmt.Println("s", s)
		fmt.Println("b", b)
	}
	cnt := 0
	const mod = 1e9 + 7
	for _, v := range b {
		cnt += v[1]
	}
	for _, v := range s {
		cnt += v[1]
	}
	return cnt % mod
}

type sellHeap [][]int
type buyHeap [][]int

func (h sellHeap) Len() int            { return len(h) }
func (h sellHeap) Less(i, j int) bool  { return h[i][0] < h[j][0] }
func (h sellHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *sellHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *sellHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func (h buyHeap) Len() int            { return len(h) }
func (h buyHeap) Less(i, j int) bool  { return h[i][0] > h[j][0] }
func (h buyHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *buyHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *buyHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
