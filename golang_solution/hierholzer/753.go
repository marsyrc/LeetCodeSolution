package hierholzer

import (
	"math"
	"strconv"
)

func crackSafe(n int, k int) string {
	seen := map[int]bool{}
	mod := int(math.Pow(10, float64(n-1)))
	ans := ""
	var dfs func(node int)
	dfs = func(node int) {
		for i := 0; i < k; i++ {
			neighbor := node*10 + i
			if !seen[neighbor] {
				seen[neighbor] = true
				dfs(neighbor % mod)
				ans += strconv.Itoa(i)
			}
		}
	}
	dfs(0)
	for i := 1; i < n; i++ {
		ans += "0"
	}
	return ans
}
