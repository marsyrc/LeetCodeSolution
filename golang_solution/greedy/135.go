package greedy

func candy(ratings []int) int {
	ans := 1
	inc, dec := 1, 0
	pre := 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] >= ratings[i-1] {
			dec = 0
			if ratings[i] == ratings[i-1] {
				pre = 1
			} else {
				pre++
			}
			ans += pre
			inc = pre
		} else {
			dec++
			if inc == dec {
				dec++
			}
			ans += dec
			pre = 1
		}
	}
	return ans
}
