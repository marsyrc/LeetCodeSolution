// 241. 为运算表达式设计优先级

// 给定一个含有数字和运算符的字符串，为表达式添加括号，改变其运算优先级以求出不同的结果。你需要给出所有可能的组合的结果。有效的运算符号包含 +, - 以及 * 。

// 示例 1:

// 输入: "2-1-1"
// 输出: [0, 2]
// 解释:
// ((2-1)-1) = 0
// (2-(1-1)) = 2

// 示例 2:

// 输入: "2*3-4*5"
// 输出: [-34, -14, -10, -10, 10]
// 解释:
// (2*(3-(4*5))) = -34
// ((2*3)-(4*5)) = -14
// ((2*(3-4))*5) = -10
// (2*((3-4)*5)) = -10
// (((2*3)-4)*5) = 10

import (
	"strconv"
)

func diffWaysToCompute(input string) []int {
	// 如果是数字，直接返回
	if isDigit(input) {
		tmp, _ := strconv.Atoi(input)
		return []int{tmp}
	}

	// 空切片
	var res []int
	// 遍历字符串
	for index, c := range input {
		tmpC := string(c)
		if tmpC == "+" || tmpC == "-" || tmpC == "*" {
			// 如果是运算符，则计算左右两边的算式
			left := diffWaysToCompute(input[:index])
			right := diffWaysToCompute(input[index+1:])

			for _, leftNum := range left {
				for _, rightNum := range right {
					var addNum int
					if tmpC == "+" {
						addNum = leftNum + rightNum
					} else if tmpC == "-" {
						addNum = leftNum - rightNum
					} else {
						addNum = leftNum * rightNum
					}
					res = append(res, addNum)
				}
			}
		}
	}

	return res
}

// 判断是否为全数字
func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	if err != nil {
		return false
	}
	return true
}
