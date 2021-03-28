package design

import (
	"strconv"
	"strings"
)

type Node struct {
	Val      int
	Children []*Node
}

type Codec struct {
}

func Constructor() *Codec {
	return &Codec{}
}

func (this *Codec) serialize(root *Node) string {
	res := []string{}
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, strconv.Itoa(node.Val))
		res = append(res, strconv.Itoa(len(node.Children)))
		for _, ch := range node.Children {
			dfs(ch)
		}
	}
	dfs(root)
	return strings.Join(res, ",")
}

func (this *Codec) deserialize(data string) *Node {
	if len(data) == 0 {
		return nil
	}
	list := strings.Split(data, ",")
	var dfs func() *Node
	dfs = func() *Node {
		root := &Node{0, []*Node{}}
		root.Val, _ = strconv.Atoi(list[0])
		list = list[1:]
		chNum, _ := strconv.Atoi(list[0])
		list = list[1:]
		root.Children = make([]*Node, chNum)
		for i := 0; i < chNum; i++ {
			root.Children[i] = dfs()
		}
		return root
	}
	return dfs()
}
