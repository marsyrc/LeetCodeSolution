/*
每次从 num 的最高位开始计算，比如58，首先得到50，
然后在整数切片(数组)中找到比50小或者相等的最大的整数，
把对应的罗马数字拼接到一个字符串中即可，
拼接之后把 num 做相应的减操作，也就是减去50。
重复上述过程，直到 num == 0 为止。
注意，在与整数比较时，不需要每次都从最大的开始比较，
只需要从上一次比较的地方开始继续比较即可。
*/
func intToRoman(num int) string {
	roman := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	intSlice := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	s := ""
	x := len(intSlice) - 1
	for num != 0{
		i := len(strconv.Itoa(num)) -1
		n := (num / int(math.Pow10(i))) * int(math.Pow10(i))
		for x>= 0{
			if n >= intSlice[x]{
				s += roman[x]
				num -= intSlice[x]
				break
			}
			x--
		}
	}
	return s
}
