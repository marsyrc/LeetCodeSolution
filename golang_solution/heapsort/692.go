package heapsort

import "container/heap"

func topKFrequent(words []string, k int) []string {
	hmap := make(map[string]int)
	for _, word := range words {
		hmap[word]++
	}
	h := IntHeap{}
	for key, v := range hmap {
		heap.Push(&h, pair{key, v})
		if len(h) > k {
			heap.Pop(&h)
		}
	}
	res := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Pop(&h).(pair).word
	}
	return res
}

type pair struct {
	word string
	cnt  int
}

type IntHeap []pair

func (h IntHeap) Len() int { return len(h) }
func (h IntHeap) Less(i, j int) bool {
	return h[i].cnt < h[j].cnt || h[i].cnt == h[j].cnt && h[i].word > h[j].word
}
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(pair)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
