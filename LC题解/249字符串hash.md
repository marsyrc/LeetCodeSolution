#### [249. 移位字符串分组](https://leetcode-cn.com/problems/group-shifted-strings/)

给定一个字符串，对该字符串可以进行 “移位” 的操作，也就是将字符串中每个字母都变为其在字母表中后续的字母，比如：`"abc" -> "bcd"`。这样，我们可以持续进行 “移位” 操作，从而生成如下移位序列：

```
"abc" -> "bcd" -> ... -> "xyz"
```

给定一个包含仅小写字母字符串的列表，将该列表中所有满足 “移位” 操作规律的组合进行分组并返回。

 

**示例：**

```
输入：["abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"]
输出：
[
  ["abc","bcd","xyz"],
  ["az","ba"],
  ["acef"],
  ["a","z"]
]
解释：可以认为字母表首尾相接，所以 'z' 的后续为 'a'，所以 ["az","ba"] 也满足 “移位” 操作规律。
```



## 方法：

寻找模式hash，可以用大素数hash匹配数值，也可以原地修改子串，这里实现了后者。

```go
func groupStrings(strings []string) [][]string {

	hashMap := make(map[string]int, len(strings))

	ans := make([][]string, 0, len(strings))

	index := 0

	for _, key := range strings {
		sr := []rune(key)
		if sr[0] != 'a' {
			for i := 1; i < len(sr); i++ {
				sr[i] = (sr[i]-sr[0]+26)%26 + 'a'
			}
			sr[0] = 'a'
		}

		ss := string(sr)
		if _, ok := hashMap[ss]; ok {
			ans[hashMap[ss]] = append(ans[hashMap[ss]], key)
			continue
		}

		hashMap[ss] = index
		index++

		ans = append(ans, []string{key})
	}
	return ans
}

```

