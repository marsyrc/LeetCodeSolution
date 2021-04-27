package mathproblem

//非进位 + 进位
//直到进位为0
func getSum(a int, b int) int {
	for b != 0 {
		a, b = a^b, (a&b)<<1
	}
	return a
}
