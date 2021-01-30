package mathproblem

func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	s := ""
	index := len(intSlice) - 1
	for num != 0 {
		for num < intSlice[index] && index > 0 {
			index--
		}
		num -= intSlice[index]
		s = s + roman[index]
	}
	return s
}
