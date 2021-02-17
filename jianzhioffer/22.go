package jianzhioffer

type ListNode struct {
	Val  int
	Next *ListNode
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow, fast := head, head
	for fast != nil && k > 0 {
		k--
		fast = fast.Next
	}
	if k != 0 {
		return nil
	}
	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
