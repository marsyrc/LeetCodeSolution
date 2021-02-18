package mathproblem

import "strconv"

func findNthDigit(n int) int {
	start, digit, count := 1, 1, 9
	for n > count {
		n -= count
		digit += 1
		start *= 10
		count = digit * start * 9
	}
	//所求数位在start开始的第(n-1)/digit个数字里面
	num := start + (n-1)/digit
	ans := strconv.Itoa(num)[(n-1)%digit] - '0'
	return int(ans)
}
