package stringproblem

import (
	"sort"
	"strings"
)

// 输入：phrases = ["writing code","code rocks"]
// 输出：["writing code rocks"]
func beforeAndAfterPuzzles(phrases []string) []string {
	n := len(phrases)
	res := []string{}
	hm := make(map[string]struct{})
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}

			//i : before, j : after
			ii := strings.LastIndexByte(phrases[i], ' ')
			ij := strings.IndexByte(phrases[j], ' ')

			if phrases[i][ii+1:] == phrases[j] || ij != -1 && phrases[i][ii+1:] == phrases[j][:ij] {
				cur := phrases[i][:ii+1] + phrases[j]
				if _, ok := hm[cur]; !ok {
					res = append(res, cur)
					hm[cur] = struct{}{}
				}
			}
		}
	}
	sort.Strings(res)
	return res
}
