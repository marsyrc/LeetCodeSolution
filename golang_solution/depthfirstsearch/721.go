package depthfirstsearch

import (
	"fmt"
	"sort"
)

func accountsMerge(accounts [][]string) [][]string {
	graph := map[string][]string{}
	for _, account := range accounts {
		master := account[1]
		for i := 2; i < len(account); i++ {
			email := account[i]
			graph[master] = append(graph[master], email)
			graph[email] = append(graph[email], master)
		}
	}
	visited := make(map[string]bool)
	var dfs func(email string, emails *[]string)
	dfs = func(email string, emails *[]string) {
		if visited[email] {
			return
		}
		visited[email] = true
		*emails = append(*emails, email)
		for _, neighbor := range graph[email] {
			dfs(neighbor, emails)
		}
	}

	res := [][]string{}
	for _, account := range accounts {
		emails := &[]string{}
		dfs(account[1], emails)
		if len(*emails) == 0 {
			continue
		}
		sort.Strings(*emails)
		fmt.Println(*emails)
		*emails = append([]string{account[0]}, *emails...)
		res = append(res, *emails)
	}
	return res
}
