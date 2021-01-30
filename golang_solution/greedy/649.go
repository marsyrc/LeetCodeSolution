package greedy

func predictPartyVictory(senate string) string {
	n := len(senate)
	qR := []int{}
	qD := []int{}
	for i := range senate {
		c := senate[i]
		if c == 'R' {
			qR = append(qR, i)
		}
		if c == 'D' {
			qD = append(qD, i)
		}
	}

	//看看当前谁能投，能投的下一轮也能投
	//不能投的和能投的都从队列里去除
	for len(qD) != 0 && len(qR) != 0 {
		if qD[0] < qR[0] {
			qD = append(qD, qD[0]+n)
		} else {
			qR = append(qR, qR[0]+n)
		}
		qD = qD[1:]
		qR = qR[1:]
	}
	if len(qD) != 0 {
		return "Dire"
	}
	return "Radiant"
}
