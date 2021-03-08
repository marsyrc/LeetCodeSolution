package stringproblem

func removeComments(source []string) []string {
	in_block := false
	res := []string{}
	newline := []byte{}
	for _, line := range source {
		cur := []byte(line)
		if !in_block {
			newline = []byte{}
		}
		i := 0
		for i < len(cur) {
			if i < len(cur)-1 && cur[i] == '/' && cur[i+1] == '*' && !in_block {
				in_block = true
				i++
			} else if i < len(cur)-1 && cur[i] == '*' && cur[i+1] == '/' && in_block {
				in_block = false
				i++
			} else if i < len(cur)-1 && cur[i] == '/' && cur[i+1] == '/' && !in_block {
				break
			} else if !in_block {
				newline = append(newline, cur[i])
			}
			i++
		}
		if len(newline) > 0 && !in_block {
			res = append(res, string(newline))
		}
	}
	return res
}
