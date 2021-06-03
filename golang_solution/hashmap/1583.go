package hashmap

func unhappyFriends(n int, preferences [][]int, pairs [][]int) int {
	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
	}

	for u, p := range preferences {
		for i, v := range p {
			memo[u][v] = i
		}
	}
	h := make(map[int]int)
	for _, pair := range pairs {
		h[pair[0]] = pair[1]
		h[pair[1]] = pair[0]
	}
	res := 0
	for x := 0; x < n; x++ {
		y := h[x]
		idx := memo[x][y]
		for i := 0; i < idx; i++ {
			u := preferences[x][i]
			v := h[u]
			if memo[u][x] < memo[u][v] {
				res++
				break
			}
		}
	}
	return res
}
