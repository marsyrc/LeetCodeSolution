package topologicalsort

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	ans := []int{}

	//记录每组负责的项目列表
	groupItems := make([][]int, m+n)
	for i := range group {
		if group[i] == -1 {
			group[i] = m + i
		}
		groupItems[group[i]] = append(groupItems[group[i]], i)
	}

	//组间拓扑
	groupGraph := make([][]int, m+n)
	groupDegree := make([]int, m+n)

	//组内拓扑
	itemGraph := make([][]int, n)
	itemDegree := make([]int, n)

	//建图
	for cur, items := range beforeItems {
		curGroupID := group[cur]
		for _, pre := range items {
			preGroupID := group[pre]
			if preGroupID != curGroupID {
				groupGraph[preGroupID] = append(groupGraph[preGroupID], curGroupID)
				groupDegree[curGroupID]++
			} else {
				itemGraph[pre] = append(itemGraph[pre], cur)
				itemDegree[cur]++
			}
		}
	}

	items := make([]int, m+n)
	for i := range items {
		items[i] = i
	}

	//获取组间序
	groupOrders := topSort(groupGraph, groupDegree, items)
	if len(groupOrders) < len(items) {
		return nil
	}

	//按组间序在组内获得项目序，输出到ans
	for _, groupId := range groupOrders {
		items := groupItems[groupId]
		orders := topSort(itemGraph, itemDegree, items)
		if len(orders) < len(items) {
			return nil
		}
		ans = append(ans, orders...)
	}
	return ans
}

func topSort(graph [][]int, indegree []int, items []int) (orders []int) {
	q := []int{}
	for _, i := range items {
		if indegree[i] == 0 {
			q = append(q, i)
		}
	}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		orders = append(orders, cur)
		for _, son := range graph[cur] {
			indegree[son]--
			if indegree[son] == 0 {
				q = append(q, son)
			}
		}
	}
	return
}
