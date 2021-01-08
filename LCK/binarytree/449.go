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
	return dfsSerialize(root, "")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	dataSlice := strings.Split(data, ",")
	for _, v := range dataSlice {
		if v != "" {
			this.list = append(this.list, v)
		}
	}
	return this.dfsDeserialize()
}

func dfsSerialize(root *TreeNode, str string) string {
	if root == nil {
		str += "null,"
	} else {
		str += strconv.Itoa(root.Val) + ","
		str = dfsSerialize(root.Left, str)
		str = dfsSerialize(root.Right, str)
	}
	return str
}

func (this *Codec) dfsDeserialize() *TreeNode {
	if this.list[0] == "null" {
		this.list = this.list[1:]
		return nil
	}
	v, _ := strconv.Atoi(this.list[0])
	this.list = this.list[1:]
	root := &TreeNode{
		Val: v,
	}
	root.Left = this.dfsDeserialize()
	root.Right = this.dfsDeserialize()
	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor()
 * deser := Constructor()
 * tree := ser.serialize(root)
 * ans := deser.deserialize(tree)
 * return ans
 */
