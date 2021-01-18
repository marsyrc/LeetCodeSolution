package listnode

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	h := nodeHeap{}
	for _, n := range lists {
		if n != nil {
			heap.Push(&h, n)
		}
	}
	dummy := &ListNode{0, nil}
	p := dummy
	for len(h) > 0 {
		cur := heap.Pop(&h).(*ListNode)
		if cur.Next != nil {
			heap.Push(&h, cur.Next)
		}
		p.Next = cur
		p = p.Next
	}
	return dummy.Next
}

type nodeHeap []*ListNode

func (h nodeHeap) Len() int            { return len(h) }
func (h nodeHeap) Less(i, j int) bool  { return h[i].Val < h[j].Val }
func (h nodeHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(x interface{}) { *h = append(*h, x.(*ListNode)) }
func (h *nodeHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}
