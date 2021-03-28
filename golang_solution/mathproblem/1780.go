package mathproblem

func checkPowersOfThree(n int) bool {
	for n != 0 {
		if n%3 == 2 {
			return false
		}
		n = n / 3
	}
	return true
}
