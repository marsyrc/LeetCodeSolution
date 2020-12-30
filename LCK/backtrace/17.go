package backtrace

func letterCombinations(digits string) (res []string) {
	if digits == "" {
		return []string{}
	}
	dict := make(map[int]string)
	dict[2] = "abc"
	dict[3] = "def"
	dict[4] = "ghi"
	dict[5] = "jkl"
	dict[6] = "mno"
	dict[7] = "pqrs"
	dict[8] = "tuv"
	dict[9] = "wxyz"

	path := []byte{}
	n := len(digits)
	var backTrace func(index int)
	backTrace = func(index int) {
		if index == n {
			res = append(res, string(path))
			return
		}
		s := dict[int(digits[index]-'0')]
		for i := 0; i < len(s); i++ {
			path = append(path, s[i])

			backTrace(index + 1)

			path = path[:len(path)-1]
		}
	}
	backTrace(0)
	return res
}
