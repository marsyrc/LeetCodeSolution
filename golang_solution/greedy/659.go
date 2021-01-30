package greedy

import "container/heap"

func isPossible(nums []int) bool {
	dict := map[int]*minHeap{}
	for _, v := range nums {
		if dict[v] == nil {
			dict[v] = new(minHeap)
		}
		if h := dict[v-1]; h != nil {
			curlen := heap.Pop(h).(int)
			if h.Len() == 0 {
				delete(dict, v-1)
			}
			heap.Push(dict[v], curlen+1)
		} else {
			heap.Push(dict[v], 1)
		}
	}

	for _, h := range dict {
		if (*h)[0] < 3 {
			return false
		}
	}
	return true
}

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
