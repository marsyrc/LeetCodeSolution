// 767. 重构字符串

// 给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。

// 若可行，输出任意可行的结果。若不可行，返回空字符串。

// 示例 1:

// 输入: S = "aab"
// 输出: "aba"

// 示例 2:

// 输入: S = "aaab"
// 输出: ""

//用大根堆，不断pop出现次数最多的元素，体现贪心。
import "container/heap"

var cnt [26]int //统计字母出现次数

type maxHeap []int

func (h maxHeap) Len() int            { return len(h) }
func (h maxHeap) Less(i, j int) bool  { return cnt[h[i]] > cnt[h[j]] } //重新实现
func (h maxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *maxHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}

func reorganizeString(S string) string {
	n := len(S)
	ss := []rune(S)
	maxcnt := 0
	cnt = [26]int{}
	for _, v := range ss {
		cnt[int(v-'a')]++
		maxcnt = max(maxcnt, cnt[int(v-'a')])
	}
	if maxcnt > (n+1)/2 {
		return ""
	}

	//constructor of maxrootheap
	h := maxHeap{}
	for i := range cnt {
		if cnt[i] > 0 {
			heap.Push(&h, i)
		}
	}

	// heap.Init(&h)
	ans := make([]byte, 0, n)
	for len(h) >= 2 {
		//一次pop出最多的两个
		i := heap.Pop(&h).(int)
		j := heap.Pop(&h).(int)
		ans = append(ans, byte('a'+i), byte('a'+j))

		//如果这两个还有就再push进去
		if cnt[i]--; cnt[i] > 0 {
			heap.Push(&h, i)
		}
		if cnt[j]--; cnt[j] > 0 {
			heap.Push(&h, j)
		}
	}

	//如果多出1个就附上
	if len(h) > 0 {
		ans = append(ans, byte('a'+h[0]))
	}
	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}