package jianzhioffer

func printNumbers(n int) []int {
	if n == 0 {
		return []int{0}
	}
	cap := 0
	for n > 0 {
		n--
		cap = cap*10 + 9
	}
	res := make([]int, cap)
	for i := 0; i < cap; i++ {
		res[i] = i + 1
	}
	return res
}
