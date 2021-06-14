package greedy

func mergeTriplets(triplets [][]int, target []int) bool {
	hi := []int{}
	c2 := 0
	for i, t := range triplets {
		if t[2] <= target[2] && t[1] <= target[1] && t[0] <= target[0] {
			if t[2] == target[2] {
				c2++
			}
			hi = append(hi, i)
		}
	}
	if c2 == 0 {
		return false
	}

	c1 := 0
	for _, idx := range hi {
		t := triplets[idx]
		if t[1] == target[1] {
			c1++
		}
	}
	if c1 == 0 {
		return false
	}

	c0 := 0
	for _, idx := range hi {
		t := triplets[idx]
		if t[0] == target[0] {
			c0++
		}
	}
	if c0 == 0 {
		return false
	}
	return true
}

//[[7,15,15],[11,8,3],[5,3,4],[12,9,9],[5,12,10],[7,15,10],[7,6,4],[3,9,8],[2,13,1],[14,2,3]]
//[14,6,4]
