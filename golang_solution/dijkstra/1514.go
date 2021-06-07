package dijkstra

import (
	"container/heap"
)

type Pair struct {
	prob float64
	id   int
}

//大根堆
type IntHeap []Pair

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i].prob > h[j].prob }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(Pair)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func maxProbability(n int, edges [][]int, succProb []float64, start int, end int) float64 {
	res := make([]float64, n)
	//建图
	graph := make([][]Pair, n)
	for i, edge := range edges {
		u, v := edge[0], edge[1]
		graph[v] = append(graph[v], Pair{succProb[i], u})
		graph[u] = append(graph[u], Pair{succProb[i], v})
	}

	//堆优化 dijkstra
	h := IntHeap{}
	heap.Push(&h, Pair{1, start})
	res[start] = 1
	for len(h) > 0 {
		top := heap.Pop(&h).(Pair)
		cur := top.id
		for _, nei := range graph[cur] {
			pNext := nei.prob
			nodeNext := nei.id
			if res[nodeNext] < res[cur]*pNext {
				res[nodeNext] = res[cur] * pNext
				heap.Push(&h, Pair{res[nodeNext], nodeNext})
			}
		}
	}
	return res[end]
}
