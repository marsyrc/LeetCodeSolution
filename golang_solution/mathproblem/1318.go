package mathproblem

func minFlips(a int, b int, c int) int {
	res := 0
	for i := 0; i < 32; i++ {
		curC := (c >> i) & 1
		sum := (a>>i)&1 + (b>>i)&1
		if curC == 0 {
			res += sum
		} else if sum == 0 {
			res += 1
		}
	}
	return res
}
