package hierholzer

import "sort"

func findItinerary(tickets [][]string) []string {
	graph := make(map[string][]string)
	for _, t := range tickets {
		graph[t[0]] = append(graph[t[0]], t[1])
	}
	for i := range graph {
		sort.Strings(graph[i])
	}
	ans := []string{}
	var dfs func(f string)
	dfs = func(f string) {
		for len(graph[f]) > 0 {
			v := graph[f][0]
			graph[f] = graph[f][1:]
			dfs(v)
		}
		ans = append([]string{f}, ans...)
	}
	dfs("JFK")
	return ans
}
