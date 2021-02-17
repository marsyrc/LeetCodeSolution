package jianzhioffer

//递归 = 栈
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	res := []int{}
	var recur func(node *ListNode)
	recur = func(node *ListNode) {
		if node == nil {
			return
		}
		recur(node.Next)
		res = append(res, node.Val)
	}
	recur(head)
	return res
}
