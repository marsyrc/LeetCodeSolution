package main

import (
	"math"
)

type edge struct {
	p      int
	weight int
}

func maxLength(arr []string) int {
	getBit := func(word string) int {
		res := 0
		for _, c := range word {
			if res&(1<<int(c-'a')) != 0 {
				return 0
			}
			res |= (1 << int(c-'a'))
		}
		return res
	}

	res := 0
	var dfs func(idx int, bit int, l int)
	dfs = func(idx int, bit int, l int) {
		if idx == len(arr) {
			res = max(res, l)
			return
		}
		for i := idx; i < len(arr); i++ {
			nextbit := getBit(arr[i])
			dfs(i+1, bit, l)
			if nextbit == 0 || nextbit&bit != 0 {
				continue
			}
			dfs(i+1, bit|nextbit, l+len(arr[i]))
		}
	}
	dfs(0, 0, 0)
	return res
}

func main() {
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	curline := input.Text()
	// }

}

/*
	Integar Utils funcs
*/
func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
	Data structure : ListNode
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

/*
	Data structure : TreeNode
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	Data structure : Heap
*/
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

/*
	Data structure : Union find set
*/
type UnionFindSet struct {
	id    []int
	sz    []int
	count int
}

func newUnionFindSet(n int) UnionFindSet {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}

	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1
	}

	return UnionFindSet{
		id:    id,
		sz:    sz,
		count: n,
	}
}

func (u *UnionFindSet) find(p int) int {
	for p != u.id[p] {
		p = u.id[p]
	}
	return p
}

func (u *UnionFindSet) union(p, q int) {
	i := u.find(p)
	j := u.find(q)
	if i == j {
		return
	}
	if u.sz[i] < u.sz[j] {
		u.id[i] = j
		u.sz[j] += u.sz[i]
	} else {
		u.id[j] = i
		u.sz[i] += u.sz[j]
	}
	u.count--
}

func (u *UnionFindSet) isConnected(p, q int) bool {
	return u.find(q) == u.find(p)
}

/*
	Data structure : Trie
*/
type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = &Trie{}
		}
		this = this.next[v-'a']
	}
	this.isEnd = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	if this.isEnd == false {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	return true
}
