package greedy

//two pointers
//i -> source, j -> target
func shortestWay(source string, target string) int {
	j, cnt := 0, 0
	for j < len(target) {
		temp := j
		i := 0
		for i < len(source) {
			if source[i] == target[j] {
				j++
				if j == len(target) {
					break
				}
			}
			i++
		}
		if temp == j {
			return -1
		}
		cnt++
	}
	return cnt
}
