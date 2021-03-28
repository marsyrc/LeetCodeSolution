package dijkstra

import (
	"container/heap"
	"math"
)

func networkDelayTime(times [][]int, n int, k int) int {
	graph := make(map[int][][]int)
	for _, e := range times {
		if _, ok := graph[e[0]]; !ok {
			graph[e[0]] = make([][]int, 0)
		}
		graph[e[0]] = append(graph[e[0]], []int{e[1], e[2]})
	}

	dist := make(map[int]int)
	h := IntHeap{}
	heap.Push(&h, []int{k, 0})
	for len(h) > 0 {
		top := heap.Pop(&h).([]int)
		d := top[1]
		node := top[0]
		if _, ok := dist[node]; ok {
			continue
		}
		dist[node] = d
		if _, ok := graph[node]; ok {
			for _, e := range graph[node] {
				nei := e[0]
				d2 := e[1]
				if _, ok2 := dist[nei]; !ok2 {
					heap.Push(&h, []int{nei, d + d2})
				}
			}
		}
	}

	if len(dist) != n {
		return -1
	}

	ans := 0
	for _, cand := range dist {
		ans = max(ans, cand)
	}
	return ans
}

type IntHeap [][]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][1] < h[j][1] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
