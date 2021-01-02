package dynamicplanning

func numDecodings(s string) int {
	n := len(s)
	str := make([]int, n)
	for i, v := range s {
		str[i] = int(v - 'a')
	}

	if n == 0 || str[0] == 0 {
		return 0
	}

	pre, cur := 1, 1
	for i := 1; i < n; i++ {
		tmp := cur
		if str[i] == 0 {
			if str[i-1] == 1 || str[i-1] == 2 {
				cur = pre
			} else {
				return 0
			}
		} else if str[i-1] == 1 || str[i-1] == 2 && str[i] >= 1 && str[i] <= 6 {
			cur += pre
		}
		pre = tmp
	}
	return cur
}
