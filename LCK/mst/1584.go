package mst

import "container/heap"

//Prim on weighted edges
func minCostConnectPoints(points [][]int) (res int) {
	visited := make(map[int]bool)
	graph := make([][]edge, len(points))
	//make graph
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			newEdge := edge{
				v:     i,
				u:     j,
				mDist: mDist(points[i][0], points[i][1], points[j][0], points[j][1]),
			}
			graph[i] = append(graph[i], newEdge)
			graph[j] = append(graph[j], newEdge)
		}
	}

	//cool func, keep thinking!
	allVisited := func(e edge) int {
		_, ok1 := visited[e.u]
		_, ok2 := visited[e.v]
		if ok1 && !ok2 {
			return e.v
		} else if ok2 && !ok1 {
			return e.u
		} else {
			return -1
		}
	}
	//maintain a PQ, storing availiable edges
	h := IntHeap{}
	visited[0] = true
	for _, e := range graph[0] {
		heap.Push(&h, e)
	}
	for len(visited) != len(points) {
		//delete invalid edges on top of PQ
		for allVisited(h[0]) == -1 {
			heap.Pop(&h)
		}

		cur := heap.Pop(&h).(edge)
		curIndex := allVisited(cur)
		visited[curIndex] = true
		res += cur.mDist
		for _, e := range graph[curIndex] {
			heap.Push(&h, e)
		}
	}
	return
}

func mDist(x1, y1, x2, y2 int) int {
	return abs(x2-x1) + abs(y2-y1)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

type edge struct {
	v     int
	u     int
	mDist int
}

type IntHeap []edge

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i].mDist < h[j].mDist }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(edge)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
