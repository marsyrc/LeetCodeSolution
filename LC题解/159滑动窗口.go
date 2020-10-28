import "fmt"

// 159. 至多包含两个不同字符的最长子串

// 给定一个字符串 s ，找出 至多 包含两个不同字符的最长子串 t ，并返回该子串的长度。

// 示例 1:

// 输入: "eceba"
// 输出: 3
// 解释: t 是 "ece"，长度为3。

// 示例 2:

// 输入: "ccaabbb"
// 输出: 5
// 解释: t 是 "aabbb"，长度为5。

func lengthOfLongestSubstringTwoDistinct(s string) int {

	left, right := 0, 0
	dict := make(map[uint8]int)
	count := 0
	ans := 0
	for right < len(s) {
		c := s[right]
		right++

		if dict[c] == 0 {
			count++
		}
		dict[c]++

		fmt.Println(left, right)

		for count > 2 {
			d := s[left]
			left++

			if dict[d] != 0 {
				dict[d]--
				if dict[d] == 0 {
					count--
				}
			}
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