package hashmap

import "sort"

func assignBikes(workers [][]int, bikes [][]int) []int {
	hm := make(map[int][]pair)
	for i, bike := range bikes {
		for j, worker := range workers {
			curDis := abs(bike[0]-worker[0]) + abs(bike[1]-worker[1])
			if _, ok := hm[curDis]; !ok {
				hm[curDis] = []pair{}
			}
			hm[curDis] = append(hm[curDis], pair{i, j})
		}
	}

	distances := []int{}
	for k := range hm {
		distances = append(distances, k)
	}
	sort.Ints(distances)

	res := make([]int, len(workers))
	usedW := make([]bool, len(workers))
	usedB := make([]bool, len(bikes))
	for _, dis := range distances {
		pairs := hm[dis]
		for _, p := range pairs {
			if !usedB[p.x] && !usedW[p.y] {
				usedB[p.x] = true
				usedW[p.y] = true

				res[p.y] = p.x
			}
		}
	}
	return res
}

type pair struct {
	x int
	y int
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
