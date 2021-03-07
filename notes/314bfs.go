import "sort"

// 314. 二叉树的垂直遍历

// 给定一个二叉树，返回其结点 垂直方向（从上到下，逐列）遍历的值。

// 如果两个结点在同一行和列，那么顺序则为 从左到右。

// 示例 1：

// 输入: [3,9,20,null,null,15,7]

//    3
//   /\
//  /  \
// 9   20
//     /\
//    /  \
//   15   7

// 输出:

// [
//   [9],
//   [3,15],
//   [20],
//   [7]
// ]

//在节点中记录当前列值
type NodeWithColumn struct {
	Column int
	Node   *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	answer := make([][]int, 0)
	if root != nil {
		m := make(map[int][]int, 0)
		verticalOrderBFS(root, &m)
		var keys []int
		for k := range m {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			answer = append(answer, m[k])
		}
	}
	return answer
}

func verticalOrderBFS(root *TreeNode, m *map[int][]int) {
	var nodes []NodeWithColumn
	nodes = append(nodes, NodeWithColumn{
		Column: 0,
		Node:   root,
	})
	for len(nodes) > 0 {
		node := nodes[0]
		nodes = nodes[1:]
		mv := *m
		mv[node.Column] = append(mv[node.Column], node.Node.Val)
		if node.Node.Left != nil {
			nodes = append(nodes, NodeWithColumn{
				Column: node.Column - 1,
				Node:   node.Node.Left,
			})
		}
		if node.Node.Right != nil {
			nodes = append(nodes, NodeWithColumn{
				Column: node.Column + 1,
				Node:   node.Node.Right,
			})
		}
	}
}
