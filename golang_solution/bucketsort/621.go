package bucketsort

import (
	"sort"
)

func leastInterval(tasks []byte, n int) int {
	lens := len(tasks)
	cnt := make([]int, 26)
	for _, v := range tasks {
		cnt[v-'A']++
	}
	sort.Slice(cnt, func(i, j int) bool {
		return cnt[i] > cnt[j]
	})
	numsOfBuckets := cnt[0]
	maxTaskCnt := 1 //最后一排的任务数量
	for i := 1; i < 26 && cnt[i] == cnt[0]; i++ {
		maxTaskCnt++
	}
	return max(lens, maxTaskCnt+(n+1)*(numsOfBuckets-1))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
