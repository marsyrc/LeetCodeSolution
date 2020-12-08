/*
此处实现的并查集不需要或者说是不应该做路径压缩
选择父亲节点应该选择层数较深的，如图中24号应该选择绿色的线，而不是红色的
如果可选的父亲节点有多个（即最深的层数相同），则随便选一个均可，如图中虚线
作者：Monologue-S
链接：https://leetcode-cn.com/problems/largest-divisible-subset/solution/cbing-cha-ji-qing-xi-tu-jie-by-monologue-s/
*/
func largestDivisibleSubset(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	father := make([]int, n) //记录父节点
	level := make([]int, n)  //记录深度

	for i := 0; i < n; i++ { //初始化
		father[i] = i
	}

	for i := 0; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if nums[i]%nums[j] == 0 && level[i] < level[j]+1 {
				father[i] = j
				level[i] = level[j] + 1
			}
		}
	}

	//get deepest element's index of level
	maxlevel := -1
	x := -1
	for i, v := range level {
		if v > maxlevel {
			x = i
			maxlevel = v
		}
	}

	ans := []int{nums[x]}
	for x != father[x] {
		x = father[x]
		ans = append(ans, nums[x])
	}
	return ans
}