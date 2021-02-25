package stringproblem

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	words := strings.Split(text, " ")
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})
	for i, w := range words {
		if w[0] >= 'A' && w[0] <= 'Z' {
			words[i] = string(byte(w[0]-'A'+'a')) + w[1:]
		}
	}
	res := strings.Join(words, " ")
	res = string(byte(res[0]-'a'+'A')) + res[1:]
	return res
}
