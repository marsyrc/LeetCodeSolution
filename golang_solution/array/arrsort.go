package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// 数组组成最大数
// 给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。
// 示例 1:
// 输入: [10,1,2]
// 输出: 2110
// 示例 2:
// 输入: [3,30,34,5,9]
// 输出: 9534330
// 作者：字节校园
// 链接：https://leetcode-cn.com/leetbook/read/bytedance-c01/eus0ut/

func main() {
	var s string
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s = input.Text()
		if s != "" {
			break
		}
	}
	str := []byte(s)
	str = str[1 : len(str)-1]
	s = string(str)
	strs := strings.Split(s, ",")
	sort.Slice(strs, func(i, j int) bool {
		return strs[i]+strs[j] > strs[j]+strs[i]
	})
	res := ""
	for _, v := range strs {
		res += v
	}
	fmt.Println(res)
}
