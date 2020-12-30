package breadthfirstsearch

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	dict := make(map[string]bool, 0)
	for _, v := range wordList {
		dict[v] = true
	}
	if !dict[endWord] {
		return [][]string{}
	}
	dict[beginWord] = true
	graph := make(map[string][]string, 0)
	distance := bfs(endWord, dict, graph)
	res := make([][]string, 0)
	dfs(beginWord, endWord, &res, []string{}, distance, graph)
	return res
}

func dfs(beginWord string, endWord string, res *[][]string, path []string, distance map[string]int, graph map[string][]string) {

	if beginWord == endWord {
		path = append(path, beginWord)
		tmp := make([]string, len(path))
		copy(tmp, path)
		(*res) = append((*res), tmp)
		path = path[:len(path)-1]
		return
	}
	for _, v := range graph[beginWord] {
		if distance[beginWord] == distance[v]+1 {
			path = append(path, beginWord)
			dfs(v, endWord, res, path, distance, graph)
			path = path[:len(path)-1]
		}
	}
}

func bfs(endWord string, dict map[string]bool, graph map[string][]string) map[string]int {
	distance := make(map[string]int, 0)
	queue := make([]string, 0)
	queue = append(queue, endWord)
	distance[endWord] = 0
	for len(queue) != 0 {
		cursize := len(queue)
		for i := 0; i < cursize; i++ {

			word := queue[0]
			queue = queue[1:]
			expansion := expand(word, dict)
			for _, v := range expansion {
				graph[v] = append(graph[v], word)
				if _, ok := distance[v]; !ok {
					distance[v] = distance[word] + 1
					queue = append(queue, v)
				}
			}
		}
	}
	return distance
}

func expand(word string, dict map[string]bool) []string {
	expansion := make([]string, 0)
	chs := []rune(word)
	for i := 0; i < len(word); i++ {
		tmp := chs[i]
		for c := 'a'; c <= 'z'; c++ {
			if tmp == c {
				continue
			}
			chs[i] = c
			newstr := string(chs)
			if dict[newstr] {
				expansion = append(expansion, newstr)
			}
		}
		chs[i] = tmp
	}
	return expansion
}
