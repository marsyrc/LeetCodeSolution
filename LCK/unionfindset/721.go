package unionfindset

import "sort"

func accountsMerge(accounts [][]string) (ans [][]string) {
	emailsToIdx := map[string]int{}
	emailsToName := map[string]string{}
	//Init
	for _, acc := range accounts {
		name := acc[0]
		for i := 1; i < len(acc); i++ {
			email := acc[i]
			if _, ok := emailsToIdx[email]; !ok {
				emailsToIdx[email] = len(emailsToIdx)
				emailsToName[email] = name
			}
		}
	}

	//union emails' idxs
	uf := newUnionFindSet(len(emailsToIdx))
	for _, acc := range accounts {
		rootIdx := emailsToIdx[acc[1]]
		for i := 2; i < len(acc); i++ {
			uf.union(rootIdx, emailsToIdx[acc[i]])
		}
	}

	//find root and record
	indexToEmails := map[int][]string{}
	for email, index := range emailsToIdx {
		index = uf.find(index)
		indexToEmails[index] = append(indexToEmails[index], email)
	}

	for _, emails := range indexToEmails {
		sort.Strings(emails)
		account := append([]string{emailsToName[emails[0]]}, emails...)
		ans = append(ans, account)
	}
	return
}

type unionFindSet struct {
	id    []int
	size  []int
	count int
}

func newUnionFindSet(n int) unionFindSet {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}

	size := make([]int, n)
	for i := range size {
		size[i] = 1
	}

	return unionFindSet{
		id:    id,
		size:  size,
		count: n,
	}
}

func (uf *unionFindSet) union(p, q int) {
	i, j := uf.find(p), uf.find(q)
	if i == j {
		return
	}
	if uf.size[i] >= uf.size[j] {
		uf.id[j] = i
		uf.size[i] += uf.size[j]
	} else {
		uf.id[i] = j
		uf.size[j] += uf.size[i]
	}
	uf.count--
}

func (uf *unionFindSet) find(p int) int {
	for p != uf.id[p] {
		p = uf.id[p]
	}
	return p
}

func (uf *unionFindSet) isConnected(p, q int) bool {
	if uf.find(p) == uf.find(q) {
		return true
	}
	return false
}
