package jianzhioffer

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	list := make([]int, 26)
	for _, value := range s {
		list[value-'a']++
	}
	for i := 0; i < len(s); i++ {
		if list[s[i]-'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
