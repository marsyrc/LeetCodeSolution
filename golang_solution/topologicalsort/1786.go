package topologicalsort

import (
	"container/heap"
	"math"
)

var graph []map[int]int

func countRestrictedPaths(n int, edges [][]int) int {
	graph = make([]map[int]int, n+1)
	dis := networkDelayTime(edges, n, n)
	dp := make([]int, n+1)
	dp[n] = 1
	minHeap := &IntHeap{}
	heap.Push(minHeap, [2]int{dis[n], n})
	dic := map[int]bool{}
	for (*minHeap).Len() > 0 {
		// 优先队列，每次都将 dis 值最小的点取出来
		x := heap.Pop(minHeap).([2]int)
		for i := range graph[x[1]] {
			if dis[i] > x[0] && dis[i] <= dis[1] {
				// 比 dis[1] 大的值没有意义，因为不可能符合受限路径条件，加==号是因为需要计算dp[1]+=
				dp[i] += dp[x[1]]
				dp[i] %= 1000000007
				if dis[i] < dis[1] && !dic[i] {
					dic[i] = true
					// 防止重复加入队列
					heap.Push(minHeap, [2]int{dis[i], i})
				}
			}
		}

	}
	return dp[1]
}

// 先套用 [743].网络延迟时间 的代码，将返回值改成 dis 切片
func networkDelayTime(times [][]int, n int, k int) []int {

	for i := range graph {
		graph[i] = map[int]int{}
	}
	for _, i := range times {
		graph[i[0]][i[1]] = i[2]
		graph[i[1]][i[0]] = i[2]
	}
	minHeap := &IntHeap{}
	dis := make([]int, n+1)
	for i := range dis {
		dis[i] = math.MaxInt32
	}
	dis[k] = 0
	for i := 1; i <= n; i++ {
		heap.Push(minHeap, [2]int{dis[i], i})
	}
	dic := map[int]bool{}
	for (*minHeap).Len() > 0 {
		x := heap.Pop(minHeap).([2]int)
		if dic[x[1]] {
			continue
		}
		dic[x[1]] = true
		for i := range graph[x[1]] {
			if dis[x[1]]+graph[x[1]][i] < dis[i] {
				dis[i] = dis[x[1]] + graph[x[1]][i]
				heap.Push(minHeap, [2]int{dis[i], i})
			}
		}
	}

	return dis
}

type IntHeap [][2]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([2]int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
