/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	window := make(map[uint8]int)
	need := make(map[uint8]int)

	//利用hashmap存储模式串遍历结果
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left, right := 0, 0
	valid := 0

	start := 0
	length := math.MaxInt32

	for right < len(s) {

		c := s[right]
		right++ //窗口右移

		if need[c] > 0 { //更新窗口内的数据
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}

		for valid == len(need) { //判断左侧窗口是否收缩
			if right-left < length {
				start = left
				length = right - left
			}

			d := s[left]
			left++ //窗口左移

			if need[d] > 0 { //更新窗口内数据
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}

	if length == math.MaxInt32 {
		return ""
	} else {
		fmt.Print(length)
		return s[start : start+length]
	}
}

// @lc code=end
