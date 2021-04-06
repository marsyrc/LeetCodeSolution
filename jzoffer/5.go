package jianzhioffer

func replaceSpace(s string) string {
	str := []byte(s)
	res := []byte{}
	for i := range str {
		if str[i] == ' ' {
			res = append(res, '%', '2', '0')
		} else {
			res = append(res, str[i])
		}
	}
	return string(res)
}
