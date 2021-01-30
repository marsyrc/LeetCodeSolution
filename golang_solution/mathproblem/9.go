package mathproblem

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}
	a := 0
	for x > a {
		a = a*10 + x%10
		x = x / 10
	}
	return x == a || x == a/10
}
