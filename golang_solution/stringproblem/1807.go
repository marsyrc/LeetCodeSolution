package stringproblem

func evaluate(s string, knowledge [][]string) string {
	hm := make(map[string][]byte)
	for _, k := range knowledge {
		hm[k[0]] = []byte(k[1])
	}
	i := 0
	res := []byte{}
	for i < len(s) {
		if s[i] == '(' {
			j := i + 1
			for j < len(s) && s[j] != ')' {
				j++
			}
			if v, ok := hm[s[i+1:j]]; ok {
				res = append(res, []byte(v)...)
			} else {
				res = append(res, '?')
			}
			i = j + 1
		} else {
			res = append(res, s[i])
			i++
		}
	}
	return string(res)
}
