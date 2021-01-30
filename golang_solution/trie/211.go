package trie

type WordDictionary struct {
	trie *Trie
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		trie: &Trie{},
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	this.trie.Insert(word)
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	return this.trie.Search(word, 0)
}

type Trie struct {
	isEnd bool
	next  [26]*Trie
}

func (t *Trie) Insert(word string) {
	c := t
	for _, v := range word {
		if c.next[v-'a'] == nil {
			c.next[v-'a'] = &Trie{}
		}
		c = c.next[v-'a']
	}
	c.isEnd = true
}

func (t *Trie) Search(word string, k int) bool {
	c := t
	for i := k; i < len(word); i++ {
		if word[i] == '.' {
			for _, v := range c.next {
				if v != nil && v.Search(word, i+1) {
					return true
				}
			}
			return false
		}
		if c.next[word[i]-'a'] == nil {
			return false
		}
		c = c.next[word[i]-'a']
	}
	return c.isEnd
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
