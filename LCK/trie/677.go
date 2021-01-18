package trie

type MapSum struct {
	val  int
	next [26]*MapSum
}

/** Initialize your data structure here. */
func Constructor() MapSum {
	return MapSum{}
}

func (this *MapSum) Insert(key string, val int) {
	for _, v := range key {
		if this.next[v-'a'] == nil {
			this.next[v-'a'] = &MapSum{}
		}
		this = this.next[v-'a']
	}
	this.val = val
}

func (this *MapSum) Sum(prefix string) int {
	for _, v := range prefix {
		if this.next[v-'a'] == nil {
			return 0
		}
		this = this.next[v-'a']
	}
	res := 0
	var dfs func(node *MapSum)
	dfs = func(node *MapSum) {
		if node == nil {
			return
		}
		if node.val != 0 {
			res += node.val
		}
		for _, v := range node.next {
			dfs(v)
		}
	}
	dfs(this)
	return res
}

/**
 * Your MapSum object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(key,val);
 * param_2 := obj.Sum(prefix);
 */
