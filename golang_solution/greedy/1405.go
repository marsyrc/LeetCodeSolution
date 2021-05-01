package greedy

import "sort"

func longestDiverseString(a int, b int, c int) string {
	list := []Char{}
	list = append(list, Char{'a', a})
	list = append(list, Char{'b', b})
	list = append(list, Char{'c', c})
	res := []byte{}
	for {
		sort.Slice(list, func(i, j int) bool {
			return list[i].cnt > list[j].cnt
		})
		cur := list[0].char
		if len(res) >= 2 && res[len(res)-1] == cur && res[len(res)-2] == cur {
			if list[1].cnt > 0 {
				cur = list[1].char
				list[1].cnt--
			} else {
				break
			}
		} else {
			if list[0].cnt > 0 {
				list[0].cnt--
			} else {
				break
			}
		}
		res = append(res, cur)
	}
	return string(res)
}

type Char struct {
	char byte
	cnt  int
}
