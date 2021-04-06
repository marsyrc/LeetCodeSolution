package jianzhioffer

func reverseLeftWords(s string, n int) string {
	str := []rune(s)
	reverse(str[:n])
	reverse(str[n:])
	reverse(str)
	return string(str)
}

func reverse(str []rune) {
	i, j := 0, len(str)-1
	for i < j {
		str[i], str[j] = str[j], str[i]
		i++
		j--
	}
}
