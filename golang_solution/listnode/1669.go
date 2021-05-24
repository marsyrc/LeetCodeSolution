package listnode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var p1 *ListNode
	var p2 *ListNode
	var tail2 *ListNode

	tail2 = list2
	for tail2.Next != nil {
		tail2 = tail2.Next
	}

	p1 = list1
	for i := 0; i < a-1; i++ {
		p1 = p1.Next
	}

	p2 = p1
	for i := 0; i < b-a+1; i++ {
		p2 = p2.Next
	}

	p1.Next = list2
	tail2.Next = p2.Next

	return list1
}
