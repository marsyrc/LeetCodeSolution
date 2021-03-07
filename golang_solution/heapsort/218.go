package heapsort

import (
	"container/heap"
	"sort"
)

func getSkyline(buildings [][]int) [][]int {
	n := len(buildings)
	points := make([]int, 0, n*2)
	for _, building := range buildings {
		points = append(points, building[0], building[1])
	}
	sort.Ints(points)
	h := buildingHeap{}
	heap.Push(&h, building{rBound: 1<<63 - 1, height: 0})
	var index int
	var out [][]int
	for _, point := range points {
		for index < n && buildings[index][0] == point {
			heap.Push(&h, building{
				rBound: buildings[index][1],
				height: buildings[index][2],
			})
			index++
		}
		for h[0].rBound <= point {
			heap.Pop(&h)
		}
		if len(out) == 0 || out[len(out)-1][1] != h[0].height {
			out = append(out, []int{point, h[0].height})
		}
	}
	return out
}

type building struct{ rBound, height int }
type buildingHeap []building

func (bh buildingHeap) Len() int            { return len(bh) }
func (bh buildingHeap) Less(i, j int) bool  { return bh[i].height > bh[j].height }
func (bh buildingHeap) Swap(i, j int)       { bh[i], bh[j] = bh[j], bh[i] }
func (bh *buildingHeap) Push(x interface{}) { *bh = append(*bh, x.(building)) }
func (bh *buildingHeap) Pop() interface{} {
	out := (*bh)[bh.Len()-1]
	*bh = (*bh)[:bh.Len()-1]
	return out
}
