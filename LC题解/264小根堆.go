import (
	"container/heap"
	"math"
)

// 264. 丑数 II

// 编写一个程序，找出第 n 个丑数。

// 丑数就是质因数只包含 2, 3, 5 的正整数。

// 示例:

// 输入: n = 10
// 输出: 12
// 解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。

// 说明:

//     1 是丑数。
//     n 不超过1690。

func nthUglyNumber(n int) int {
	uglyNum := -1 // 首先把丑数初始化为一个不合法的值-1
	var h IntHeap // 然后定义最小堆
	heap.Init(&h)
	heap.Push(&h, 1) // 并且把第一个丑数1，加入到最小堆中
	for n > 0 {      // 当n大于0时执行以下操作
		for h[0] == uglyNum { // 检查堆顶元素是否等于上一个丑数
			heap.Pop(&h) // 如果是就不断丢弃
		} // 这一步的作用是去重
		// 直到不重复就取出堆顶元素,把它作为新的丑数
		uglyNum = heap.Pop(&h).(int)

		if 2*uglyNum <= math.MaxInt32 {
			heap.Push(&h, 2*uglyNum)
		}
		if 3*uglyNum <= math.MaxInt32 {
			heap.Push(&h, 3*uglyNum)
		}
		if 5*uglyNum <= math.MaxInt32 {
			heap.Push(&h, 5*uglyNum)
		}
		n--
	}
	return uglyNum
}

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
  