package mathproblem

import "strings"

func numberToWords(num int) string {
	to19 := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Eleven", "Twelve",
		"Thirteen", "Fourteen", "Fifteen", "Sixteen", "Seventeen", "Eighteen", "Nineteen"}
	tens := []string{"Twenty", "Thirty", "Forty", "Fifty", "Sixty", "Seventy", "Eighty", "Ninety"}
	res := []string{}
	var recur func(n int)
	recur = func(n int) {
		if n == 0 {
			return
		}
		if n < 20 {
			res = append(res, to19[n-1])
			return
		}
		if n < 100 {
			if n%10 == 0 {
				res = append(res, tens[n/10-2])
				return
			}
			res = append(res, tens[n/10-2])
			recur(n % 10)
			return
		}
		if n < 1000 {
			res = append(res, to19[n/100-1], "Hundred")
			recur(n % 100)
			return
		}
		if n < 1000*1000 {
			recur(n / 1000)
			res = append(res, "Thousand")
			recur(n % 1000)
			return
		}
		if n < 1000*1000*1000 {
			recur(n / 1000 / 1000)
			res = append(res, "Million")
			recur(n % (1000 * 1000))
			return
		}
		if n < 1000*1000*1000*1000 {
			recur(n / 1000 / 1000 / 1000)
			res = append(res, "Billion")
			recur(n % (1000 * 1000 * 1000))
			return
		}
		return
	}
	if num == 0 {
		return "Zero"
	}
	recur(num)
	return strings.Join(res, " ")
}
