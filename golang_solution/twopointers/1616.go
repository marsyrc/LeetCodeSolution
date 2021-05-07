package twopointers

func checkPalindromeFormation(a string, b string) bool {
	if check(a, b) || check(b, a) {
		return true
	}
	a = string(reverse([]byte(a)))
	b = string(reverse([]byte(b)))
	if check(a, b) || check(b, a) {
		return true
	}
	return false
}

func check(a string, b string) bool {
	n := len(a)
	flag := true
	for i := 0; i < n/2; i++ {
		if flag {
			if a[i] != b[n-1-i] {
				flag = false
			}
		}
		if !flag {
			if a[i] != a[n-1-i] {
				return false
			}
		}
	}
	return true
}

func reverse(a []byte) []byte {
	l, r := 0, len(a)-1
	for l < r {
		a[l], a[r] = a[r], a[l]
		l++
		r--
	}
	return a
}
