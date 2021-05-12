package topologicalsort

func alienOrder(words []string) string {
	graph := make(map[byte]map[byte]struct{})
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < len(words[i]); j++ {
			if j >= len(words[i+1]) { //an illegal case
				return ""
			}
			if words[i][j] == words[i+1][j] {
				continue
			}
			if _, ok := graph[words[i][j]]; !ok {
				graph[words[i][j]] = make(map[byte]struct{})
			}
			graph[words[i][j]][words[i+1][j]] = struct{}{}
			break
		}
	}

	indegeree := make([]int, 26)
	for i := range indegeree {
		indegeree[i] = -1
	}

	for _, word := range words {
		for i := range word {
			indegeree[int(word[i]-'a')] = 0
		}
	}
	for from := range graph {
		for to := range graph[from] {
			indegeree[int(to-'a')]++
		}
	}
	res := []byte{}
	q := []byte{}
	count := 0
	for i := 0; i < 26; i++ {
		if indegeree[i] != -1 {
			count++
		}
		if indegeree[i] == 0 {
			q = append(q, byte('a'+i))
		}
	}

	//bfs
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			res = append(res, cur)
			if _, ok := graph[cur]; !ok {
				continue
			}
			for to, _ := range graph[cur] {
				indegeree[to-'a']--
				if indegeree[to-'a'] == 0 {
					q = append(q, to)
				}
			}
		}
	}
	if len(res) != count {
		return ""
	}
	return string(res)
}
