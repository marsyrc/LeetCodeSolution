package breadthfirstsearch

func ladderLength(beginWord string, endWord string, wordList []string) int {
	//init a dict of wordList
	dict := make(map[string]bool)
	for _, v := range wordList {
		dict[v] = true
	}
	//special case
	if _, ok := dict[endWord]; !ok {
		return 0
	}

	//bfs, with queue
	queue := []string{beginWord}
	step := 0 //record current levels
	for len(queue) > 0 {
		step++ //incr by
		size := len(queue)
		for i := 0; i < size; i++ {
			s := queue[0]
			queue = queue[1:]
			chs := []rune(s)
			for i, v := range chs {
				tmp := v

				for c := 'a'; c <= 'z'; c++ {
					if c == tmp {
						continue
					}
					chs[i] = c
					t := string(chs)
					//has found the terminal
					if t == endWord {
						return step + 1
					}
					if _, ok := dict[t]; !ok {
						continue
					}
					//avoid access this word twice
					delete(dict, t)
					queue = append(queue, t)
				}
				chs[i] = tmp
			}
		}
	}
	return 0
}
