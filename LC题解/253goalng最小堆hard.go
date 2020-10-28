import (
	"container/heap"
	"sort"
)

// 253. 会议室 II

// 给定一个会议时间安排的数组，每个会议时间都会包括开始和结束的时间 [[s1,e1],[s2,e2],...] (si < ei)，为避免会议冲突，同时要考虑充分利用会议室资源，请你计算至少需要多少间会议室，才能满足这些会议安排。

// 示例 1:

// 输入: [[0, 30],[5, 10],[15, 20]]
// 输出: 2

// 示例 2:

// 输入: [[7,10],[2,4]]
// 输出: 1

func minMeetingRooms(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	// 先按照开始时间从小到大排序
	scs := schedules(intervals)
	sort.Sort(scs)
	pq := make(PriorityQueue, 0, len(scs))
	heap.Push(&pq, &Item{
		value:    scs[0][1],
		priority: scs[0][1],
		index:    0,
	})

	for i := 1; i < len(scs); i++ {
		if scs[i][0] >= pq[0].value {
			heap.Pop(&pq)
		}
		heap.Push(&pq, &Item{
			value:    scs[i][1],
			priority: scs[i][1],
			index:    i,
		})
	}
	return pq.Len()
}

type schedules [][]int

func (s schedules) Len() int {
	return len(s)
}

func (s schedules) Less(i int, j int) bool {
	return s[i][0] < s[j][0]
}

func (s schedules) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

// Item 是优先队列中包含的元素。
type Item struct {
	value    int // 结束时间
	priority int // 元素在队列中的优先级。
	// 元素的索引可以用于更新操作，它由 heap.Interface 定义的方法维护。
	index int // 元素在堆中的索引。
}

// 一个实现了 heap.Interface 接口的优先队列，队列中包含任意多个 Item 结构。
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 我们希望 Pop 返回的是最小值
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // 为了安全性考虑而做的设置
	*pq = old[0 : n-1]
	return item
}
