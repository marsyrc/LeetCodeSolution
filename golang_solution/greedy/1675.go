package greedy

import (
	"container/heap"
	"math"
)

/*可以考虑把所有数都化为最大的形式(全部化为所能变成的最大偶数)，
然后加入大根堆，维护 mi 表示当前堆的最小值，堆顶即为最大值。
为了不断缩小最大值与最小值的差，
我们必须不断取出当前堆最大的数，然后将其除以2，重新加入队列。
这个过程同时维护答案 res 和 mi
当堆顶为奇数时，意味着不能再除以2，最大值不可能再改变，
差距也就不会再被缩小，此时的 res即为答案。
*/
func minimumDeviation(nums []int) int {
	h := maxHeap{}
	minNum := math.MaxInt32
	for _, v := range nums {
		if v%2 == 1 {
			heap.Push(&h, 2*v)
			minNum = min(minNum, v*2)
		} else {
			heap.Push(&h, v)
			minNum = min(minNum, v)
		}
	}
	res := math.MaxInt32
	for {
		cur := heap.Pop(&h).(int)
		res = min(res, cur-minNum)
		if cur%2 == 1 {
			break
		}
		minNum = min(minNum, cur>>1)
		heap.Push(&h, cur>>1)
	}
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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
