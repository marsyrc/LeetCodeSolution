package dynamicplanning

import "container/heap"

func kthLargestValue(matrix [][]int, k int) int {
	m, n := len(matrix), len(matrix[0])
	for j := 1; j < n; j++ {
		matrix[0][j] ^= matrix[0][j-1]
	}
	for i := 1; i < m; i++ {
		matrix[i][0] ^= matrix[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] ^= matrix[i-1][j] ^ matrix[i][j-1] ^ matrix[i-1][j-1]
		}
	}
	h := IntHeap{}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			heap.Push(&h, matrix[i][j])
			if len(h) > k {
				heap.Pop(&h)
			}
		}
	}
	return heap.Pop(&h).(int)
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
