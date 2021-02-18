package jianzhioffer

func sumNums(n int) int {
	if n <= 1 {
		return n
	} else {
		return n + sumNums(n-1)
	}
}
