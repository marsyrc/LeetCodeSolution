package main

/*给定一个用字符串展示的化学公式，计算每种元素的个数。
规则如下：
    元素命名采用驼峰命名，例如 Mg
    () 代表内部的基团，代表阴离子团
    [] 代表模内部链节通过化学键的连接，并聚合
例如：H2O => H2O1 Mg(OH)2 => H2Mg1O2
输入描述:
化学公式的字符串表达式，例如：K4[ON(SO3)2]2
输出描述:
元素名称及个数：K4N2O14S4，并且按照元素名称升序排列
测试用例
输入
K4[ON(SO3)2]2
输出
K4N2O14S4
*/
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		s = input.Text()
		if s != "" {
			break
		}
	}
	res := ""
	count = make(map[string]int)
	solution(0, len(s)-1, 1)
	keys := []string{}
	for key, _ := range count {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		cur := key + strconv.Itoa(count[key])
		res += cur
	}
	fmt.Println(res)
}

var s string
var count map[string]int

//每遇到一个左括号就进入递归
func solution(l, r, cnt int) {
	for i := l; i <= r; i++ {
		t := string(s[i])
		if isBig(s[i]) {
			if i+1 <= r && isDigit(s[i+1]) {
				count[t] += cnt * int(s[i+1]-'0')
			} else if i+1 <= r && isSmall(s[i+1]) {
				t += string(s[i+1])
				if i+2 <= r && isDigit(s[i+2]) {
					count[t] += cnt * int(s[i+2]-'0')
				} else {
					count[t] += cnt
				}
				i++
			} else {
				count[t] += cnt
			}
		} else if s[i] == '[' || s[i] == '(' {
			var c byte
			if s[i] == '[' {
				c = ']'
			} else {
				c = ')'
			}
			temp := i + 1
			k := 1
			for s[i+1] != c {
				i++
			}
			if i+2 <= r && isDigit(s[i+2]) {
				k = int(s[i+2] - '0')
			}
			solution(temp, i, k*cnt)
		}
	}
}

func isBig(b byte) bool {
	return b >= 'A' && b <= 'Z'
}

func isSmall(b byte) bool {
	return b >= 'a' && b <= 'z'
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}
