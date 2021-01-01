package dynamicplanning

import "fmt"

// //using prefix
// func maxSubarraySumCircular(A []int) int {
// 	n := len(A)
// 	P := make([]int, n * 2 + 1)
// 	for i := 0; i < n * 2; i++ {
// 		P[i + 1]= P[i] + A[i % n]
// 	}
// 	res := A[0]
// 	q := []int{}
// 	q = append(q, 0)
// 	for j := 1; j <= 2 * n; j++ {
// 		if j - q[0] > n {
// 			q = q[1:]
// 		}
// 		res = max(res, P[j] - P[q[0]])

// 		for len(q) != 0 && P[j] <= P[q[len(q) - 1]] {
// 			q = q[:len(q) - 1]
// 		}
// 		q = append(q, j)
// 	}
// 	return res
// }

//在环形交界处取得最大等于在A两头取得最小
//在53.go基础上，实现一个最小子序和，让两头联结起来最大，这个最小子序和考虑范围为[1, n - 2]
func maxSubarraySumCircular(A []int) int {
	n := len(A)
	if n <= 2 {
		a, _ := maxSubArray(A)
		return a
	}
	A1, A2 := make([]int, n), make([]int, n)
	copy(A1, A)
	copy(A2, A)
	res1, sum := maxSubArray(A1)
	res2 := sum - minSubArray(A2)
	ans := max(res1, res2)
	return ans
}

func maxSubArray(nums []int) (int, int) {
	sum := nums[0]
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		sum += nums[i]
		if nums[i-1] > 0 {
			nums[i] += nums[i-1]
		}
		res = max(res, nums[i])
	}
	return res, sum
}

func minSubArray(nums []int) int {
	res := nums[1]
	fmt.Println(res)
	for i := 2; i < len(nums)-1; i++ {
		if nums[i-1] < 0 {
			nums[i] += nums[i-1]
		}
		res = min(res, nums[i])
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
