package breadthfirstsearch

func findLexSmallestString(s string, a int, b int) string {
	res := s
	q := []string{}
	q = append(q, res)
	vis := make(map[string]struct{})
	for len(q) > 0 {
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			vis[cur] = struct{}{}
			if cur < res {
				res = cur
			}
			n1 := acc(cur, a)
			n2 := mov(cur, b)
			if _, ok1 := vis[n1]; !ok1 {
				vis[n1] = struct{}{}
				q = append(q, n1)
			}
			if _, ok2 := vis[n2]; !ok2 {
				vis[n2] = struct{}{}
				q = append(q, n2)
			}
		}
	}
	return res
}

func acc(s string, a int) string {
	ss := make([]int, len(s))
	res := []byte{}
	for i, c := range s {
		ss[i] = int(c - '0')
		if i&1 == 1 {
			ss[i] = (ss[i] + a) % 10
		}
		res = append(res, byte(ss[i]+'0'))
	}
	return string(res)
}

func mov(s string, b int) string {
	return s[len(s)-b:] + s[:len(s)-b]
}
