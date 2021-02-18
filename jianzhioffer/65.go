package jianzhioffer

func add(a int, b int) int {
	carry := 0
	for b != 0 {
		carry = (a & b) << 1
		a ^= b
		b = carry
	}
	return a
}
