package stack

func removeDuplicates(S string) string {
	stack := []byte{}
	for i := range S {
		ch := S[i]
		if len(stack) > 0 && stack[len(stack)-1] == ch {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}
	return string(stack)
}
