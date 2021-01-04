package dynamicplanning

import "math"

//single string with dimension K,
//try to fix K, make it be single string problem.

//dp[i][k] : [i, n - 1] divided into K packets

//let [i, j] has been made into one packets,
//so [j, n - 1] should be divided into K - 1 packets
//ergodic every j of range [i, n - 1]
//then dp[i][k] = max of dp[j][k - 1] + sum(i , j) / (j - 1)
func largestSumOfAverages(A []int, K int) float64 {
	n := len(A)

	//use prefix array to calculate sum(i, j)
	P := make([]int, n+1)
	for i := 1; i < n; i++ {
		P[i] = A[i] + P[i-1]
	}

	//init : k = 0
	dp := make([]float64, n)
	for i := 0; i < n; i++ {
		dp[i] = float64(P[n]-P[i]) / float64(n-i)
	}

	for k := 1; k < K; k++ {
		for i := 0; i < n; i++ {
			for j := i + 1; j < n; j++ {
				dp[i] = math.Max(dp[i], dp[j]+float64(P[j]-P[i])/float64(j-i))
			}
		}
	}
	return dp[0]
}
