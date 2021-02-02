package twopointers

//贪心思想 + 双指针
//i -> source, j -> target
//在source中找以i开头的最长适配j的序列，如果j没变说明没找到，返回-1
func shortestWay(source string, target string) int {
	j := 0
	cnt := 0
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
		if j == temp {
			return -1
		}
		cnt++
	}
	return cnt
}
