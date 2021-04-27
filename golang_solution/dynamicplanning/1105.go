package dynamicplanning

func minHeightShelves(books [][]int, shelf_width int) int {
	n := len(books)
	dp := make([]int, n+1)
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 1000 * 1000
	}

	for i := 1; i <= n; i++ {
		curWidth := 0
		j := i
		h := 0

		for j > 0 {
			curWidth += books[j-1][0]
			if curWidth > shelf_width {
				break
			}
			h = max(h, books[j-1][1])
			dp[i] = min(dp[i], dp[j-1]+h)
			j--
		}
	}
	return dp[n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
