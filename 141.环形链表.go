/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]int)
	for head != nil {
		_, ok := nodeMap[head]
		if ok {
			return true
		} else {
			nodeMap[head] = 0
		}
		head = head.Next
	}
	return false
}

// @lc code=end
