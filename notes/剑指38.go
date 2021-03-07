// 剑指 Offer 38. 字符串的排列
// 输入一个字符串，打印出该字符串中字符的所有排列。
// 你可以以任意顺序返回这个字符串数组，但里面不能有重复元素。
// 示例:
// 输入：s = "abc"
// 输出：["abc","acb","bac","bca","cab","cba"]
// 限制：
// 1 <= s 的长度 <= 8

//重复方案与剪枝： 当字符串存在重复字符时，排列方案中也存在重复方案。为排除重复方案，需在固定某位字符时，
//保证 “每种字符只在此位固定一次” ，即遇到重复字符时不交换，直接跳过。从 DFS 角度看，此操作称为 “剪枝” 。

func permutation(s string) []string {
	var res []string
	helper([]byte(s), 0, &res)
	return res
}

func helper(s []byte, start int, res *[]string) {
	if start == len(s) {
		*res = append(*res, string(s))
	}
	m := make(map[byte]int)
	for i := start; i < len(s); i++ {
		if _, ok := m[s[i]]; ok {
			continue
		}
		s[i], s[start] = s[start], s[i]
		helper(s, start+1, res)
		s[i], s[start] = s[start], s[i]
		m[s[i]] = 1
	}
}
  