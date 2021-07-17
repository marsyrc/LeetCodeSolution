package mathproblem

//逻辑推导
func sumGame(num string) bool {
	n := len(num)
	n0, n1, q0, q1 := 0, 0, 0, 0
	for i := 0; i < n / 2; i++ {
			if num[i] == '?' {
					q0++
			} else {
					n0 += int(num[i] - '0')
			}
	}
	for i := n / 2; i < n; i++ {
			if num[i] == '?' {
					q1++
			} else {
					n1 += int(num[i] - '0')
			}
	}
	return ((q0 + q1) % 2 == 1) || (n0 - n1 != (q1 - q0) * 9 / 2);
}
