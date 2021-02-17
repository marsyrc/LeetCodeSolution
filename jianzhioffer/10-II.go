package jianzhioffer

func numWays(n int) int {
	if n <= 1 {
		return 1
	}
	pre := 1
	cur := 1
	for i := 2; i < n+1; i++ {
		tmp := cur
		cur = (cur + pre) % 1000000007
		pre = tmp
	}
	return cur
}
