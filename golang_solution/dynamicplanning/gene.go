package main

/**DNA 是由 ACGT 四种核苷酸组成，例如 AAAGTCTGAC，假定自然环境下 DNA 发生异变的情况有：
    基因缺失一个核苷酸
    基因新增一个核苷酸
    基因替换一个核苷酸
且发生概率相同。
古生物学家 Sam 得到了若干条相似 DNA 序列，
Sam 认为一个 DNA 序列向另外一个 DNA 序列转变所需的最小异变情况数可以代表其物种血缘相近程度，
异变情况数越少，血缘越相近，请帮助 Sam 实现获取两条 DNA 序列的最小异变情况数的算法。
输入描述:
    每个样例只有一行，两个 DNA 序列字符串以英文逗号“,”分割
输出描述:
输出转变所需的最少情况数，类型是数字
测试用例:
输入
ACT,AGCT
输出
1
数据范围：
    每个 DNA 序列不超过 100 个字符
*/

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	var s string
	for input.Scan() {
		s = input.Text()
		if s != "" {
			break
		}
	}
	ss := strings.Split(s, ",")
	a, b := ss[0], ss[1]
	fmt.Println(diff(a, b))
}

func diff(a, b string) int {
	m, n := len(a), len(b)
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i][j-1]) + 1
			}
		}
	}
	return dp[m][n]
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
