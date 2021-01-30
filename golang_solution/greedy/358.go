package greedy

import "container/heap"

type maxHeap []int

var cnt [26]int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return cnt[h[i]] > cnt[h[j]] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rearrangeString(s string, k int) string {
	if k == 0 {
		return s
	}
	//计数
	ss := []byte(s)
	maxCnt := 0
	cnt = [26]int{}
	for _, c := range ss {
		cnt[int(c-'a')]++
		maxCnt = max(maxCnt, cnt[int(c-'a')])
	}
	//入堆
	h := maxHeap{}
	for i, v := range cnt {
		if v > 0 {
			heap.Push(&h, i)
		}
	}
	ans := []byte{}
	q := []int{}
	for len(h) > 0 {
		c := heap.Pop(&h).(int)
		cnt[c]--
		q = append(q, c)
		ans = append(ans, byte(c+'a'))
		if len(q) == k {
			d := q[0]
			q = q[1:]
			if cnt[d] > 0 {
				heap.Push(&h, d)
			}
		}
	}
	if len(ans) != len(s) {
		return ""
	}
	return string(ans)
}
