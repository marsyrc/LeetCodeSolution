#### [291. 单词规律 II](https://leetcode-cn.com/problems/word-pattern-ii/)

给你一种规律 `pattern` 和一个字符串 `str`，请你判断 `str` 是否遵循其相同的规律。

这里我们指的是 **完全遵循**，例如 `pattern` 里的每个字母和字符串 `str` 中每个 **非空** 单词之间，存在着 **双射** 的对应规律。**双射** 意味着映射双方一一对应，不会存在两个字符映射到同一个字符串，也不会存在一个字符分别映射到两个不同的字符串。

 

**示例 1：**

```
输入：pattern = "abab", s = "redblueredblue"
输出：true
解释：一种可能的映射如下：
'a' -> "red"
'b' -> "blue"
```

**示例 2：**

```
输入：pattern = "aaaa", s = "asdasdasdasd"
输出：true
解释：一种可能的映射如下：
'a' -> "asd"
```

## 回溯法+两个hashmap

```go
func wordPatternMatch(pattern string, s string) bool {
	//用map建立模式字符和子串之间的映射
	hm := make(map[byte]string)
	hm_t := make(map[string]byte)
	//index是在pattern中的索引，k是在字符串中的索引
	var backtrace func(index int, k int) bool
	backtrace = func(index int, k int) bool {
		//end case 
		if index == len(pattern) && k == len(s) {
			return true
		}
		if index == len(pattern) || k == len(s) {
			return false
		}

		//当前模式字符
		c := pattern[index]
		//如果能找到，则先判断用当前模式映射是否合法
		if v, ok := hm[c]; ok {
			if k+len(v) > len(s) {
				return false
			}
			if s[k:k+len(v)] != v {
				return false
			}
            //向前回溯
			return backtrace(index+1, k+len(v))
		}
        
        //如果没找到，开始循环所有的选择，并做选择、回溯
		for i := k; i < len(s); i++ {
			substr := s[k : i+1]
			if _, ok := hm_t[substr]; ok {
				continue
			}
			hm[c] = substr
			hm_t[substr] = c
			if backtrace(index+1, i+1) {
				return true
			}
			delete(hm, c)
			delete(hm_t, substr)
		}
		return false
	}
	return backtrace(0, 0)
}

```

