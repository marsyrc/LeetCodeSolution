package binarytree

type Node struct {
	Val      int
	Children []*Node
}

func findRoot(tree []*Node) *Node {
	if len(tree) == 0 || tree[0] == nil {
		return nil
	}

	xorSum := 0
	for _, node := range tree {
		xorSum ^= node.Val
		for _, childNode := range node.Children {
			xorSum ^= childNode.Val
		}
	}

	for _, node := range tree {
		if xorSum == node.Val {
			return node
		}
	}
	return nil
}
