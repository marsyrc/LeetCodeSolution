#### [973. 最接近原点的 K 个点](https://leetcode-cn.com/problems/k-closest-points-to-origin/)

我们有一个由平面上的点组成的列表 `points`。需要从中找出 `K` 个距离原点 `(0, 0)` 最近的点。

（这里，平面上两点之间的距离是欧几里德距离。）

你可以按任何顺序返回答案。除了点坐标的顺序之外，答案确保是唯一的。

 

**示例 1：**

```
输入：points = [[1,3],[-2,2]], K = 1
输出：[[-2,2]]
解释： 
(1, 3) 和原点之间的距离为 sqrt(10)，
(-2, 2) 和原点之间的距离为 sqrt(8)，
由于 sqrt(8) < sqrt(10)，(-2, 2) 离原点更近。
我们只需要距离原点最近的 K = 1 个点，所以答案就是 [[-2,2]]。
```

**示例 2：**

```
输入：points = [[3,3],[5,-1],[-2,4]], K = 2
输出：[[3,3],[-2,4]]
（答案 [[-2,4],[3,3]] 也会被接受。）
```

### 解法一直接排序获得呢（失去offer！）

```go
import (
	"math"
	"sort"
)

func kClosest(points [][]int, K int) [][]int {
	if K == len(points) {
		return points
	}
	var dis func(point []int) float64
	dis = func(point []int) float64{
		return math.Sqrt(float64(point[0]) * float64(point[0]) + float64(point[1]) * float64(point[1]))
	}
	sort.Slice(points, func(i, j int) bool {
		return dis(points[i]) > dis(points[j])
	})
	return points[len(points) - K: ]
}

```

### 解法二利用大根heap（Less方法中决定是小根还是大根堆）

把最大的根放在堆顶，比当前堆顶小就替换，并再平衡

用Fix 比 Pop再Push的方法更快呢

```go
import "container/heap"

type pair struct {
	dist   int
	points []int
}

type hp []pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h hp) Less (i, j int) bool {
	return h[i].dist > h[j].dist
}

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(pair))
}

func (h *hp) Pop() interface{} {
	tmp := *h
	n := len(tmp)
	v := tmp[n-1]
	*h = tmp[:n-1]
	return v
}

func kClosest(points [][]int, K int) (ans [][]int) {
	h := make(hp, K)
	for i, p := range points[:K] {
		h[i] = pair{p[0]*p[0] + p[1]*p[1], p}
	}
	heap.Init(&h)
	for _, p := range points[K:]{
		if dist := p[0] * p[0] + p[1] * p[1]; dist < h[0].dist {
			h[0] = pair{dist, p}
			heap.Fix(&h, 0)
		}
	}
	for _, p := range h {
		ans = append(ans, p.points)
	}
	return ans
}

```

