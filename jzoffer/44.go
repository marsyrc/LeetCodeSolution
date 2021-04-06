package jianzhioffer

import "strconv"

func findNthDigit(n int) int {
	if n == 0 {
		return 0
	}
	n -= 1
	start := 1
	digits := 1
	cnt := 9
	for n > cnt {
		n -= cnt
		start *= 10
		digits++
		cnt = start * digits * 9
	}
	num := start + n/digits
	ans := strconv.Itoa(num)[n%digits] - '0'
	return int(ans)

}
