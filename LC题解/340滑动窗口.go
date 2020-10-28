// 340. 至多包含 K 个不同字符的最长子串

// 给定一个字符串 s ，找出 至多 包含 k 个不同字符的最长子串 T。

// 示例 1:

// 输入: s = "eceba", k = 2
// 输出: 3
// 解释: 则 T 为 "ece"，所以长度为 3。

// 示例 2:

// 输入: s = "aa", k = 1
// 输出: 2
// 解释: 则 T 为 "aa"，所以长度为 2。

func lengthOfLongestSubstringKDistinct(s string, k int) int {
	if k == 0 {
		return 0
	}
	seen := map[byte]int{}
	cnt := 0
	ans := 0

	left := 0
	right := 0

	for right < len(s) {
		c := s[right]
		if seen[c] == 0 {
			cnt++
		}
		seen[c]++
		right++

		for cnt > k {
			d := s[left]
			seen[d]--
			if seen[d] == 0 {
				cnt--
			}
			left++
		}

		ans = max(ans, right-left)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}