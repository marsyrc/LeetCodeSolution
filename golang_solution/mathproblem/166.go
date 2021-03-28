package mathproblem

import "strconv"

func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}
	res := []byte{}
	if numerator*denominator < 0 {
		res = append(res, '-')
	}
	dividend := abs(numerator)
	divisor := abs(denominator)
	res = append(res, []byte(strconv.Itoa(dividend/divisor))...)
	remainder := dividend % divisor
	if remainder == 0 {
		return string(res)
	}
	hmap := make(map[int]int)
	fraction := []byte{}
	i := 0
	for remainder != 0 {
		if idx, ok := hmap[remainder]; ok {
			temp := []byte{}
			temp = append(temp, fraction[:idx]...)
			temp = append(temp, '(')
			temp = append(temp, fraction[idx:i]...)
			temp = append(temp, ')')
			return string(res) + "." + string(temp)
		}
		hmap[remainder] = i
		remainder *= 10
		fraction = append(fraction, []byte(strconv.Itoa(remainder/divisor))...)
		remainder %= divisor
		i++
	}
	return string(res) + "." + string(fraction)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
