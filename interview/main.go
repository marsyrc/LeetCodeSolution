package main

import (
	"errors"
	"math"
)

type edge struct {
	p      int
	weight int
}

func minSideJumps(obstacles []int) int {
	dp1, dp2, dp3 := 1, 0, 1
	for i := 0; i < len(obstacles); i++ {
		switch obstacles[i] {
		case 1:
			if i > 0 && obstacles[i-1] == 2 {
				dp2 = dp3 + 1
			} else {
				dp2 = min(dp2, dp1+1)
			}
			if i > 0 && obstacles[i-1] == 3 {
				dp3 = dp2 + 1
			} else {
				dp3 = min(dp3, dp1+1)
			}
			dp1 += 2
		case 2:
			if i > 0 && obstacles[i-1] == 1 {
				dp1 = dp3 + 1
			} else {
				dp1 = min(dp1, dp2+1)
			}
			if i > 0 && obstacles[i-1] == 3 {
				dp3 = dp1 + 1
			} else {
				dp3 = min(dp3, dp2+1)
			}
			dp2 += 2
		case 3:
			if i > 0 && obstacles[i-1] == 1 {
				dp1 = dp2 + 1
			} else {
				dp1 = min(dp1, dp3+1)
			}
			if i > 0 && obstacles[i-1] == 2 {
				dp2 = dp1 + 1
			} else {
				dp2 = min(dp2, dp3+1)
			}
			dp3 += 2
		}
	}
	n := len(obstacles)
	if obstacles[n-1] == 1 {
		return min(dp2, dp3)
	} else if obstacles[n-1] == 2 {
		return min(dp1, dp3)
	} else if obstacles[n-1] == 3 {
		return min(dp1, dp2)
	}
	return min(dp1, dp2, dp3)
}

func main() {
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	curline := input.Text()
	// }

}

/*
	Data structure : ListNode
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func solution(year int, month int, day int) (int, error) {
	var isLeapYear bool
	isLeapYear = year%4 == 0 && year%100 != 0 || year%400 == 0
	dayInMonths := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if isLeapYear {
		dayInMonths[1] = 29
	}

	var err error
	err = errors.New("invalid input!")
	if month <= 0 || month > 12 || day <= 0 || day > dayInMonths[month-1] {
		return 0, err
	}

	res := 0
	for i := 0; i < month-1; i++ {
		res += dayInMonths[i]
	}
	res += day
	return res, nil
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
