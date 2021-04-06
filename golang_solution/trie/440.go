package trie

func findKthNumber(n int, k int) int {
	getCount := func(prefix int) int {
		cur, next := prefix, prefix+1
		count := 0
		for cur <= n {
			// 下一个前缀的起点减去当前前缀的起点就是 prefix 开始到 n 之间的本层节点数量
			count += min(next, n+1) - cur
			cur *= 10
			next *= 10
		}
		return count
	}
	preNum := 1
	parent := 1
	for preNum < k {
		count := getCount(parent)
		if count+preNum > k {
			parent *= 10
			preNum++
		} else {
			parent++
			preNum += count
		}
	}
	return parent
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
