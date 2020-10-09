/*
187. 重复的DNA序列

所有 DNA 都由一系列缩写为 A，C，G 和 T 的核苷酸组成，例如：“ACGAATTCCG”。在研究 DNA 时，识别 DNA 中的重复序列有时会对研究非常有帮助。

编写一个函数来查找目标子串，目标子串的长度为 10，且在 DNA 字符串 s 中出现次数超过一次。



示例：

输入：s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
输出：["AAAAACCCCC", "CCCCCAAAAA"]

*/

//hashmap去重尚可，可用rabin-karp算法，ACTG四进制映射

func findRepeatedDnaSequences(s string) []string {
	n := 10
	l := len(s)
	dict := make(map[string]int, l-n+1)
	for i := 0; i < l-n+1; i++ {
		dict[s[i:i+10]]++
	}
	ans := []string{}
	for k, v := range dict {
		if v > 1 {
			ans = append(ans, k)
		}
	}
	return ans
}