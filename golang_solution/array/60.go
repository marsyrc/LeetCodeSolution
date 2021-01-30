package array

import "strconv"

func getPermutation(n int, k int) string {
	factorial := make([]int, n+1) //阶乘
	factorial[0] = 1

	//tokens 描述从1到n的整数集合
	tokens := make([]int, n)
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
		tokens[i-1] = i
	}

	k--
	var b string
	for n > 0 {
		n--
		a := k / factorial[n]
		k = k % factorial[n]

		b += strconv.Itoa(tokens[a])
		tokens = append(tokens[:a], tokens[a+1:]...)
	}
	return b
}
