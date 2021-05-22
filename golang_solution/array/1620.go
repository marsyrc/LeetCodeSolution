package array

import "math"

func bestCoordinate(towers [][]int, radius int) []int {
	u, r := 51, 51
	d, l := 0, 0
	count := func(i, j int) int {
		sigs := 0
		for _, t := range towers {
			x, y := t[0], t[1]
			d := (x-i)*(x-i) + (y-j)*(y-j)
			sig := 0
			if d <= radius*radius {
				sig = int(float64(t[2]) / (1 + math.Sqrt(float64(d))))
			}
			sigs += sig
		}
		return sigs
	}
	maxSigs := 0
	res := []int{0, 0}
	for i := l; i <= r; i++ {
		for j := d; j <= u; j++ {
			cur := count(i, j)
			if maxSigs < cur {
				res = []int{i, j}
				maxSigs = cur
			}
		}
	}
	return res
}

func min(a ...int) int {
	res := math.MaxInt32
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}

func max(a ...int) int {
	res := math.MinInt32
	for _, v := range a {
		if v > res {
			res = v
		}
	}
	return res
}
