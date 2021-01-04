package prefixsum

type NumMatrix struct {
	prefix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		prefix := make([][]int, 1)
		for i := range prefix {
			prefix[i] = make([]int, 1)
		}
		return NumMatrix{
			prefix: prefix,
		}
	} else {
		m, n := len(matrix), len(matrix[0])
		dp := make([][]int, m+1)
		for i := range dp {
			dp[i] = make([]int, n+1)
		}
		for i := 1; i <= m; i++ {
			for j := 1; j <= n; j++ {
				dp[i][j] = dp[i-1][j] + dp[i][j-1] - dp[i-1][j-1] + matrix[i-1][j-1]
			}
		}
		return NumMatrix{
			prefix: dp,
		}
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.prefix[row2+1][col2+1] + this.prefix[row1][col1] - this.prefix[row1][col2+1] - this.prefix[row2+1][col1]
}
