package greedy

func canCompleteCircuit(gas []int, cost []int) int {
	start := 0
	sum, box := 0, 0
	for i := range gas {
		sum += gas[i] - cost[i]
		box += gas[i] - cost[i]
		if box >= 0 {
			continue
		}
		//不够，只能从下一个开始
		start = i + 1
		box = 0
	}
	if sum < 0 {
		return -1
	}
	return start
}
