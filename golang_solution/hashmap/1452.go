package hashmap

import "sort"

//优化的暴力比较
func isSub(sub, par []string) bool {
	lenSub, lenPar := len(sub), len(par)
	pSub, pPar := 0, 0
	for pSub < lenSub && pPar < lenPar {
		if sub[pSub] == par[pPar] {
			pSub++
			pPar++
		} else {
			pPar++
		}
	}
	return pSub == lenSub
}

func peopleIndexes(favoriteCompanies [][]string) []int {
	for _, com := range favoriteCompanies {
		sort.Strings(com)
	}
	sorted := make([][]string, len(favoriteCompanies))
	copy(sorted, favoriteCompanies)
	sort.SliceStable(sorted, func(i, j int) bool {
		return len(sorted[i]) >= len(sorted[j])
	})
	var ans = make([]int, 0, len(favoriteCompanies))
	for i, sub := range favoriteCompanies {
		ans = append(ans, i)
		for _, par := range sorted {
			if len(sub) >= len(par) {
				break
			} else if isSub(sub, par) {
				ans = ans[:len(ans)-1]
				break
			}
		}
	}
	return ans
}
