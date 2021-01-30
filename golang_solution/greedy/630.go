package greedy

import (
	"container/heap"
	"sort"
)

func scheduleCourse(courses [][]int) (res int) {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	cur := 0
	pq := maxHeap{}
	for _, v := range courses {
		if cur+v[0] <= v[1] {
			heap.Push(&pq, v[0])
			cur += v[0]
		} else if len(pq) != 0 && pq[0] > v[0] {
			cur += v[0] - heap.Pop(&pq).(int)
			heap.Push(&pq, v[0])
		}
	}
	return len(pq)
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
