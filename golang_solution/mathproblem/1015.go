package mathproblem

/*
Given a positive integer K,
you need to find the length of the smallest positive integer N such that N is divisible by K,
and N only contains the digit 1.
Return the length of N. If there is no such N, return -1.
Note: N may not fit in a 64-bit signed integer.
*/

func smallestRepunitDivByK(K int) int {
	length := 1
	x := 1
	m := map[int]bool{}
	for {
		x %= K
		if x == 0 {
			break
		}
		if _, seen := m[x]; seen {
			return -1
		}
		m[x] = true
		x = x*10 + 1
		length++
	}
	return length
}

func smallestRepunitDivByK2(K int) int {
	x := 1
	length := 1
	if K%2 == 0 || K%5 == 0 {
		return -1
	}
	for x%K != 0 {
		x = x % K
		x = x*10 + 1
		length++
	}
	return length
}
