package slidingwindow

// 使用一个队列保存 i 前面 k - 1 个位置有多少元素被反转了。
// 如果队列长度为奇数，那么当前位置的1被变成 0 了需要反转，
// 如果为偶数，说明当前位置的 0 还是 0，需要反转。
// 如果最后 k - 1 个位置还有0的话说明失败。否则将 i 加入队列，更新答案。
func minKBitFlips(A []int, K int) int {
	q := []int{}
	res := 0
	n := len(A)
	for i := 0; i < n; i++ {
		for len(q) > 0 && q[0]+K <= i {
			q = q[1:]
		}
		//当前位在前面的翻转后为0，需要反转
		if len(q)%2 == A[i] {
			if i+K > n {
				return -1
			}
			q = append(q, i)
			res++
		}
	}
	return res
}
