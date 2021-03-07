package trie

import (
	"sort"
)

func suggestedProducts(products []string, searchWord string) [][]string {
	tr := Constructor()
	sort.Strings(products)
	for _, p := range products {
		tr.Insert(p)
	}
	return tr.Search(searchWord)
}

type Trie struct {
	isEnd bool
	next  [26]*Trie
	cache []string
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	for _, v := range word {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = new(Trie)
		}
		if len(this.next[v-'a'].cache) < 3 {
			this.next[v-'a'].cache = append(this.next[v-'a'].cache, word)
		}
		// fmt.Println(this.cache)
		this = this.next[v-'a']
	}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) [][]string {
	res := [][]string{}
	flag := -1
	for i, v := range word {
		if this.next[v-'a'] == nil {
			res = append(res, []string(nil))
			flag = i
			break
		} else {
			tmp := make([]string, len(this.next[v-'a'].cache))
			copy(tmp, this.next[v-'a'].cache)
			res = append(res, tmp)
			this = this.next[v-'a']
		}
	}
	if flag != -1 {
		for i := 0; i < len(word)-flag-1; i++ {
			res = append(res, []string(nil))
		}
	}
	return res
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
