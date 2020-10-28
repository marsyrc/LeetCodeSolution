func furthestBuilding(heights []int, bricks int, ladders int) int {
	count := 0
	var f func(pos int, bricks int, ladders int, heights []int)
	f = func(pos int, bricks int, ladders int, heights []int) {
		if count == len(heights)-1 {
			return
		}

		if pos == (len(heights)-1) || (heights[pos+1]-heights[pos] > bricks && ladders == 0) {
			count = max(count, pos)
			return
		}

		if heights[pos+1] <= heights[pos] {
			f(pos+1, bricks, ladders, heights)
		} else {
			dff := heights[pos+1] - heights[pos]
			if dff <= bricks {
				f(pos+1, bricks-dff, ladders, heights)
			}
			if ladders >= 1 {
				f(pos+1, bricks, ladders-1, heights)
			}
		}
	}
	f(0, bricks, ladders, heights)
	return count
}
