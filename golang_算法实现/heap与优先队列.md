## 导包

```go
import (
	"container/heap"
)
```

## 需要实现的接口

```
type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1
}
```

sort.Interface中的内容：

```
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

需要对任意数据类型实现以上五个方法，以使用heap接口

## 利用方式

```go
h := &IntHeap{3, 8, 6}  // 创建IntHeap类型的原始数据
func Init(h Interface)  // 对heap进行初始化，生成小根堆（或大根堆）
func Push(h Interface, x interface{})  // 往堆里面插入内容
func Pop(h Interface) interface{}  // 从堆顶pop出内容
func Remove(h Interface, i int) interface{}  // 从指定位置删除数据，并返回删除的数据
func Fix(h Interface, i int)  // 从i位置数据发生改变后，对堆再平衡，优先级队列使用到了该方法


//小根堆

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
  
```

## 利用heap实现优先队列

```go
import (
	"container/heap"
	"fmt"
)

type Item struct {
	value    string // 优先级队列中的数据，可以是任意类型，这里使用string
	priority int    // 优先级队列中节点的优先级
	index    int    // index是该节点在堆中的位置
}

// 优先级队列需要实现heap的interface
type PriorityQueue []*Item

// 绑定Len方法
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// 绑定Less方法，这里用的是小于号，生成的是小根堆
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

// 绑定swap方法
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index, pq[j].index = i, j
}

// 绑定put方法，将index置为-1是为了标识该数据已经出了优先级队列了
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	item.index = -1
	return item
}

// 绑定push方法
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// 更新修改了优先级和值的item在优先级队列中的位置
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	// 创建节点并设计他们的优先级
	items := map[string]int{"二毛": 5, "张三": 3, "狗蛋": 9}
	i := 0
	pq := make(PriorityQueue, len(items)) // 创建优先级队列，并初始化
	for k, v := range items {             // 将节点放到优先级队列中
		pq[i] = &Item{
			value:    k,
			priority: v,
			index:    i}
		i++
	}
	heap.Init(&pq) // 初始化堆
	item := &Item{ // 创建一个item
		value:    "李四",
		priority: 1,
	}
	heap.Push(&pq, item)           // 入优先级队列
	pq.update(item, item.value, 6) // 更新item的优先级
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("%.2d:%s index:%.2d\n", item.priority, item.value, item.index)
	}
}
```

### 可以做leetcode973

## leetcode296

实现了一个大根堆和小根堆,平衡两个队，能取到中位数

```go
import "container/heap"

type minHeap []int

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *minHeap) Pop() interface{} {
   res := (*h)[len(*h)-1]
   *h = (*h)[:len(*h)-1]
   return res
}

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
   res := (*h)[len(*h)-1]
   *h = (*h)[:len(*h)-1]
   return res
}

type MedianFinder struct {
   maxh *maxHeap
   minh *minHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
   res := MedianFinder{
      maxh: new(maxHeap),
      minh: new(minHeap),
   }

   return res
}

func (this *MedianFinder) AddNum(num int) {
   heap.Push(this.maxh, num)
   heap.Push(this.minh, (heap.Pop(this.maxh)).(int))

   if len(*this.maxh) < len(*this.minh) {
      heap.Push(this.maxh, (heap.Pop(this.minh)).(int))
   }
}

func (this *MedianFinder) FindMedian() float64 {
   if len(*this.minh) == len(*this.maxh) {
      return (float64((*this.maxh)[0]) + float64((*this.minh)[0])) / float64(2)
   } else {
      return float64((*this.maxh)[0])
   }
}
```

