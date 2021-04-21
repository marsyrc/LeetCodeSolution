package breadthfirstsearch

import "sort"

func watchedVideosByFriends(watchedVideos [][]string, friends [][]int, id int, level int) []string {
	cnt := make(map[string]int)
	//build friendGraph
	friendGraph := make(map[int][]int)
	for i, v := range friends {
		_, ok := friendGraph[i]
		if !ok {
			friendGraph[i] = []int{}
		}
		for _, friend := range v {
			friendGraph[i] = append(friendGraph[i], friend)
		}
	}

	visited := make([]bool, len(watchedVideos))
	q := []int{id}
	visited[id] = true
	i := 0
	for len(q) > 0 {
		if i == level {
			break
		}
		sz := len(q)
		for i := 0; i < sz; i++ {
			cur := q[0]
			q = q[1:]
			for _, friend := range friendGraph[cur] {
				if !visited[friend] {
					visited[friend] = true
					q = append(q, friend)
				}
			}
		}
		i++
	}
	for _, people := range q {
		for _, video := range watchedVideos[people] {
			cnt[video]++
		}
	}
	res := []string{}
	for video := range cnt {
		res = append(res, video)
	}
	sort.Slice(res, func(i, j int) bool {
		return cnt[res[i]] < cnt[res[j]] || cnt[res[i]] == cnt[res[j]] && res[i] < res[j]
	})
	return res
}
