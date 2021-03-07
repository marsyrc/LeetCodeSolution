package binarytree

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
	list []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	var dfsS func(node *TreeNode, str string) string
	dfsS = func(node *TreeNode, str string) string {
		if node == nil {
			str += "null,"
			return str
		}
		str += strconv.Itoa(node.Val) + ","
		str = dfsS(node.Left, str)
		str = dfsS(node.Right, str)
		return str
	}
	return dfsS(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	this.list = strings.Split(data, ",")
	var dfsD func() *TreeNode
	dfsD = func() *TreeNode {
		if this.list[0] == "null" {
			this.list = this.list[1:]
			return nil
		}
		v, _ := strconv.Atoi(this.list[0])
		this.list = this.list[1:]
		root := &TreeNode{v, nil, nil}
		root.Left = dfsD()
		root.Right = dfsD()
		return root
	}
	return dfsD()
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
