// 373. 查找和最小的K对数字

// 给定两个以升序排列的整形数组 nums1 和 nums2, 以及一个整数 k。

// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2。

// 找到和最小的 k 对数字 (u1,v1), (u2,v2) ... (uk,vk)。

// 示例 1:

// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [1,2],[1,4],[1,6]
// 解释: 返回序列中的前 3 对数：
//      [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

// 示例 2:

// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [1,1],[1,1]
// 解释: 返回序列中的前 2 对数：
//      [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

// 示例 3:

// 输入: nums1 = [1,2], nums2 = [3], k = 3
// 输出: [1,3],[2,3]
// 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := IntHeap{}
	if len(nums2) == 0 || len(nums1) == 0 {
		return [][]int{}
	}
	//init
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			curList := []int{nums1[i], nums2[j]}
			heap.Push(&h, curList)
		}
	}
	res := [][]int{}
	for k > 0 {
		if len(h) == 0 {
			return res
		}
		u := heap.Pop(&h).([]int)
		res = append(res, u)
		k--
	}
	return res
}

type IntHeap [][]int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i][0]+h[i][1] < h[j][0]+h[j][1] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.([]int)) }
func (h *IntHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}