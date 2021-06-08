package hashmap

import "strconv"

func differByOne(dict []string) bool {
	hset := make(map[string]struct{})
	for _, word := range dict {
		for i := 0; i < len(word); i++ {
			cur := word[:i] + word[i+1:] + strconv.Itoa(i)
			if _, ok := hset[cur]; ok {
				return true
			}
			hset[cur] = struct{}{}
		}
	}
	return false
}
