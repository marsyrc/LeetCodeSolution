package trie

func findWords(board [][]byte, words []string) []string {
	tr := &Trie{}
	memo := map[string]int{}
	for _, word := range words {
		tr.Insert(word)
		memo[word] = 0
	}
	m, n := len(board), len(board[0])
	inArea := func(i, j int) bool {
		return i >= 0 && j >= 0 && i < m && j < n
	}
	var dfs func(i, j int, node *Trie, str string)
	dfs = func(i, j int, node *Trie, str string) {
		if !inArea(i, j) {
			return
		}
		c := board[i][j]
		if c == '#' || node.next[c-'a'] == nil {
			return
		}
		node = node.next[c-'a']
		str += string(c)
		if _, ok := memo[str]; ok {
			memo[str]++
		}
		board[i][j] = '#'
		dfs(i, j+1, node, str)
		dfs(i-1, j, node, str)
		dfs(i+1, j, node, str)
		dfs(i, j-1, node, str)
		board[i][j] = c
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			p := tr
			dfs(i, j, p, "")
		}
	}
	ans := []string{}
	for k, v := range memo {
		if v > 0 {
			ans = append(ans, k)
		}
	}
	return ans
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = &Trie{}
		}
		this = this.next[v-'a']
	}
	this.isEnd = true
}

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	if this.isEnd == false {
		return false
	}
	return true
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	for _, v := range prefix {
		if this.next[v-'a'] == nil {
			return false
		}
		this = this.next[v-'a']
	}
	return true
}
