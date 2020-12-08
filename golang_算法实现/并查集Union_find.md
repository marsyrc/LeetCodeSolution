并查集Union find，用来表示网络节点之间是否连接的集合。
 这里的网络是一抽抽象的概念，不仅包括互联网，人与人形成的网络，道路之间形成的网络，迷宫网络等等。
 数据存储格式如下

![img](C:\Users\mars\Documents\keepCoding\leetCodingSon\golang_算法实现\aHR0cHM6Ly91cGxvYWQtaW1hZ2VzLmppYW5zaHUuaW8vdXBsb2FkX2ltYWdlcy8xNTIwMzU2NS1mYjE2Y2U3YWM2ZjFiNjVjLmpwZw)

1）i 为节点索引， set[i] 为所属集合，若两个索引所属集合相同，表示这两节点相连
 2）跟传统树结构不同的是，并查集是一种子节点指向父节点的数据结构

```go
//对并查集来说，主要支持两个动作：
1）union（p，q） // 将p，q两个节点相连
2）isConnected（p，q）  // 查询p，q是否相连，即是否属于同一个集合
```

```go
import "gitee.com/kklt1996/data-structure/common"

/*
	并查集
*/
type UF interface {
	GetSize() int

	IsConnected(p int, q int) bool

	UnionElements(p int, q int) error
}

/*
	创建大小为size的并查集合
*/
func CreateUnionQuickFind(size int) *UnionQuickFind {
	id := make([]int, size, size)
	// 初始化的时候每个索引都不是联通的
	for i, _ := range id {
		id[i] = i
	}
	return &UnionQuickFind{id: id}
}

/*
	创建大小为size的快速合并和查询的并查集
*/
func CreateQuickUnionFind(size int) *QuickUnionFind {
	parent := make([]int, size, size)
	// 初始化的时候每个索引的根节点都是自己
	for i, _ := range parent {
		parent[i] = i
	}
	// 初始化的时候每个树的高度都为1,都是本身
	rank := make([]int, size, size)
	for i, _ := range rank {
		rank[i] = 1
	}
	return &QuickUnionFind{parent: parent, rank: rank}
}

/*
	并查集的实现
	基于数组实现的并查集,支持快速查询
	使用数组的索引表示元素,数组的值表示元素所属的集合id
*/
type UnionQuickFind struct {
	id []int
}

func (u UnionQuickFind) GetSize() int {
	return len(u.id)
}

/*
	O(1)
	检查数组索引代表的元素是否连通
*/
func (u UnionQuickFind) IsConnected(p int, q int) bool {
	firstSet, err := u.find(p)
	if err != nil {
		return false
	}
	secondSet, err := u.find(q)
	if err != nil {
		return false
	}
	// 检查数组索引所代表的元素是否在一个集合
	return firstSet == secondSet
}

/*
	O(n)
	将索引所代表的元素合并到一个集合
	将p和q索引所属的集合的元素都合并到一个集合中
*/
func (u *UnionQuickFind) UnionElements(p int, q int) error {
	firstSet, err := u.find(p)
	if err != nil {
		return err
	}
	secondSet, err := u.find(q)
	if err != nil {
		return err
	}
	// 集合相同本身相同就不用合并了
	if firstSet == secondSet {
		return nil
	}
	// 将两个集合中的元素归入一个集合
	for i, v := range u.id {
		if v == secondSet {
			u.id[i] = firstSet
		}
	}
	return nil
}

/*
	查询第index个节点所属的集合
*/
func (u UnionQuickFind) find(index int) (int, error) {
	if index < 0 || index >= len(u.id) {
		return -1, common.IndexError{}
	}
	return u.id[index], nil
}

/*
	快速合并和查询
	使用一种数组表示的森林来实现并查集.
	1.相同集合的元素所属同一个根
	2.合并元素的时候将一个元素所属集合的根节点修改为另一个元素所属集合的根
*/
type QuickUnionFind struct {
	// 数组的值表示索引的父亲节点的索引
	parent []int

	// 数组的值表示以索引为根的树的高度的一个参考值
	rank []int
}

func (uf QuickUnionFind) GetSize() int {
	return len(uf.parent)
}

/*
	获取集合数量
*/
func (uf QuickUnionFind) GetSetNum() int {
	var num int
	for k, v := range uf.parent {
		if k == v {
			num++
		}
	}
	return num
}

/*
	O(h) h为树的高度
	查询两个节点是否连接
*/
func (uf QuickUnionFind) IsConnected(p int, q int) bool {
	firstSet, err := uf.find(p)
	if err != nil {
		return false
	}
	secondSet, err := uf.find(q)
	if err != nil {
		return false
	}
	return secondSet == firstSet
}

/*
	O(h) h为树的高度
	将两个元素所在的集合合并成一个集合
*/
func (uf *QuickUnionFind) UnionElements(p int, q int) error {
	firstSet, err := uf.find(p)
	if err != nil {
		return err
	}
	secondSet, err := uf.find(q)
	if err != nil {
		return err
	}

	if uf.rank[secondSet] < uf.rank[firstSet] {
		// 第二个集合树的高度小于第一个集合,将第二个集合合并到第一个集合
		uf.parent[secondSet] = firstSet
	} else if uf.rank[firstSet] < uf.rank[secondSet] {
		uf.parent[firstSet] = secondSet
	} else {
		uf.parent[secondSet] = firstSet
		// 两棵树高度相等的话,合并之后高度+1
		uf.rank[firstSet] += 1
	}
	return nil
}

/*
	O(h)
	寻找index索引位置元素的根节点
*/
func (uf QuickUnionFind) find(index int) (int, error) {
	if index < 0 || index >= len(uf.parent) {
		return -1, common.IndexError{}
	}
	for index != uf.parent[index] {
		// 在查询当前元素的所属集合的时候,对查询路径进行压缩
		// 将当前节点的父亲节点,修改为父亲节点的父亲节点
		uf.parent[index] = uf.parent[uf.parent[index]]
		index = uf.parent[index]
	}
	return index, nil
}

```

