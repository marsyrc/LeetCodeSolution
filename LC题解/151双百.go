import "strings"

// 151. 翻转字符串里的单词

// 给定一个字符串，逐个翻转字符串中的每个单词。

// 说明：

//     无空格字符构成一个 单词 。
//     输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
//     如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

// 示例 1：

// 输入："the sky is blue"
// 输出："blue is sky the"

// 示例 2：

// 输入："  hello world!  "
// 输出："world! hello"
// 解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

// 示例 3：

// 输入："a good   example"
// 输出："example good a"
// 解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

// 示例 4：

// 输入：s = "  Bob    Loves  Alice   "
// 输出："Alice Loves Bob"

// 示例 5：

// 输入：s = "Alice does not even like bob"
// 输出："bob like even not does Alice"

// 提示：

//     1 <= s.length <= 104
//     s 包含英文大小写字母、数字和空格 ' '
//     s 中 至少存在一个 单词

// 进阶：

//     请尝试使用 O(1) 额外空间复杂度的原地解法。

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	s = PreProcess(s)
	wordArray := strings.Split(s, " ")
	for _, v := range wordArray {
		v = strings.TrimSpace(v)
	}
	left, right := 0, len(wordArray)-1
	for left < right {
		wordArray[left], wordArray[right] = wordArray[right], wordArray[left]
		left++
		right--
	}
	ans := strings.Join(wordArray, " ")
	return ans
}
func PreProcess(s string) string {
	l := len(s)
	var res []byte
	flag := 1                     // 用于处理多个连续空格
	for l != 0 && s[l-1] == ' ' { // 处理字符串后面的空格
		l--
	}
	for i := 0; i < l; i++ {
		if s[i] != ' ' {
			res = append(res, s[i])
			flag = 0
		}
		if s[i] == ' ' && flag == 0 {
			res = append(res, s[i])
			flag = 1
		}
	}
	return string(res)
}
