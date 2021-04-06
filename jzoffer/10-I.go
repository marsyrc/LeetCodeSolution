package jianzhioffer

func fib(n int) int {
	pre := 0
	cur := 1
	for i := 0; i < n; i++ {
		tmp := cur
		cur = (cur + pre) % 1000000007
		pre = tmp
	}
	return pre % 1000000007
}
