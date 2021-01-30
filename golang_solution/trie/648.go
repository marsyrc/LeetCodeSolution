package trie

import "strings"

func replaceWords(dict []string, sentence string) string {
	// 先创建前缀树
	trie := Trie{}
	for _, v := range dict {
		trie.Insert(v)
	}

	words := strings.Split(sentence, " ")
	for i, v := range words {
		if r := trie.FindShortestPrefix(v); r != -1 {
			words[i] = v[:r]
		}
	}

	return strings.Join(words, " ")
}

type Trie struct {
	next [26]*Trie
	end  bool
}

/** Inserts a word into the trie. */
func (t *Trie) Insert(word string) {
	current := t
	for _, ch := range word {
		index := ch - 'a'
		if current.next[index] == nil {
			current.next[index] = &Trie{}
		}
		current = current.next[index]
	}
	current.end = true
}

// 查找最短前缀，如 bat 是 battle 的前缀，前提是 bat 被添加过
// i 的初始值应为0，结束时 i 对应最短前缀的长度
func (t *Trie) FindShortestPrefix(word string) int {
	node, i := t, 0
	for _, v := range word {
		if node.end {
			return i
		}
		offset := v - 'a'
		if node.next[offset] == nil {
			return -1 // 未找到
		}
		node = node.next[offset]
		i++
	}
	return i
}
