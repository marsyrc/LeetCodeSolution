import "strings"

// 140. 单词拆分 II

// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。

// 说明：

//     分隔时可以重复使用字典中的单词。
//     你可以假设字典中没有重复的单词。

// 示例 1：

// 输入:
// s = "catsanddog"
// wordDict = ["cat", "cats", "and", "sand", "dog"]
// 输出:
// [
//   "cats and dog",
//   "cat sand dog"
// ]

// 示例 2：

// 输入:
// s = "pineapplepenapple"
// wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
// 输出:
// [
//   "pine apple pen apple",
//   "pineapple pen apple",
//   "pine applepen apple"
// ]
// 解释: 注意你可以重复使用字典中的单词。

// 示例 3：

// 输入:
// s = "catsandog"
// wordDict = ["cats", "dog", "sand", "and", "cat"]
// 输出:
// []

/**
先用单词拆分I中的dp判断当前s能否拆分，再标准回溯呢
*/
func wordBreak(s string, wordDict []string) []string {
	wordMap := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		wordMap[v] = true
	}

	//DP判断能否拆分
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordMap[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	res := []string{}
	if !dp[len(s)] {
		return res
	}

	path := []string{}
	var dfs func(chS string)
	dfs = func(chS string) {
		if len(chS) == 0 {
			res = append(res, strings.Join(path, " "))
			return
		}
		for i := 1; i <= len(chS); i++ {
			if wordMap[chS[:i]] {
				path = append(path, chS[:i])
				dfs(chS[i:])
				path = path[:len(path)-1]
			}
		}
	}
	dfs(s)
	return res
}
