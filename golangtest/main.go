package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"reflect"
	"sort"
	"sync"
	"time"
)

func maxBuilding(n int, r [][]int) int {
	r = append(r, []int{1, 0})
	sort.Slice(r, func(i, j int) bool {
		return r[i][0] < r[j][0]
	})
	if r[len(r)-1][0] != n {
		r = append(r, []int{n, n - 1})
	}
	l := len(r)
	for i := 1; i < l; i++ {
		r[i][1] = min(r[i][1], r[i-1][1]+(r[i][0]-r[i-1][0]))
	}
	for i := l - 2; i >= 0; i-- {
		r[i][1] = min(r[i][1], r[i+1][1]+(r[i+1][0]-r[i][0]))
	}

	res := -1
	for i := 1; i < l; i++ {
		res = max(res, (r[i][1]+r[i-1][1]+r[i][0]-r[i-1][0])/2)
	}
	return res
}

func Add(a, b int) int {
	return a + b
}

func main() {
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	curline := input.Text()
	// }
	v := reflect.ValueOf(Add)
	if v.Kind() != reflect.Func {
		return
	}
	t := v.Type()
	argv := make([]reflect.Value, t.NumIn())
	for i := range argv {
		if t.In(i).Kind() != reflect.Int {
			return
		}
		argv[i] = reflect.ValueOf(i)
	}
	result := v.Call(argv)
	if len(result) != 1 || result[0].Kind() != reflect.Int {
		return
	}
	fmt.Println(result[0].Int())

	ctx, cancelF := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancelF()

	mu := sync.RWMutex{}
	mu.Lock()
	
	mu.RLock()
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
