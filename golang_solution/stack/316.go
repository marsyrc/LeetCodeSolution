package stack

//贪心单调栈
func removeDuplicateLetters(s string) string {
	dict := make(map[byte]int)
	ss := []byte(s)
	for _, c := range ss {
		dict[c]++
	}

	n := len(ss)
	stack := []byte{}
	seen := make(map[byte]bool)
	for i := 0; i < n; i++ {
		c := ss[i]
		if seen[c] {
			dict[c]--
			continue
		}
		for len(stack) > 0 && dict[stack[len(stack)-1]] > 0 && stack[len(stack)-1] > c {
			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			seen[pop] = false
		}
		seen[c] = true
		dict[c]--
		stack = append(stack, c)
	}
	return string(stack)
}
