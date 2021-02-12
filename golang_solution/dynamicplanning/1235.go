package dynamicplanning

import (
	"math"
	"sort"
)

func jobScheduling(startTime []int, endTime []int, profit []int) int {
	n := len(startTime)
	jobs := make([][3]int, n)
	for i := range startTime {
		jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
	}
	//sort the jobs slice
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][1] < jobs[j][1]
	})
	dp := make([]int, n)
	dp[0] = jobs[0][2]
	maxPre := dp[0]
	for i := 1; i < n; i++ {
		j := i - 1
		for j >= 0 {
			if jobs[j][1] <= jobs[i][0] {
				dp[i] = max(dp[i], dp[j]+jobs[i][2])
				break
			}
			j--
		}
		dp[i] = max(dp[i], maxPre)
		dp[i] = max(dp[i], jobs[i][2])
		maxPre = max(maxPre, dp[i])
	}
	return dp[n-1]
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}

// 二分
// func jobScheduling(startTime []int, endTime []int, profit []int) int {
//     n := len(startTime)
//     jobs := make([][3]int, n)
//     for i := range startTime {
//         jobs[i] = [3]int{startTime[i], endTime[i], profit[i]}
//     }
//     //sort the jobs slice
//     sort.Slice(jobs, func(i, j int) bool {
//         return jobs[i][1] < jobs[j][1]
//     })
//     dp := make([]int, n)
//     dp[0] = jobs[0][2]
//     pre := 0
//     for i := 1; i < n; i++ {
//         pre = getIdx(jobs, i) - 1
//         if pre >= 0 {
//             dp[i] = max(dp[i - 1], dp[pre] + jobs[i][2])
//         } else {
//             dp[i] = max(dp[i - 1], jobs[i][2])
//         }
//     }
//     return dp[n - 1]
// }

// func getIdx(jobs [][3]int, cur int) int {
//     l, r := 0, cur
//     mid := 0
//     for l < r {
//         mid = l + (r-l) / 2
//         if jobs[mid][1] > jobs[cur][0] {
//             r = mid
//         } else {
//             l = mid + 1
//         }
//     }
//     return l
// }
// func max(a ...int) int {
// 	res := math.MinInt32
// 	for _, v := range a {
// 		if v > res {
// 			res = v
// 		}
// 	}
// 	return res
// }
