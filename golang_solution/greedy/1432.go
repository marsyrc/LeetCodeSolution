package greedy

import "math"

func maxDiff(num int) int {
	stack := []int{}
	for num != 0 {
		a := num % 10
		stack = append([]int{a}, stack...)
		num /= 10
	}
	n := len(stack)
	a, b := make([]int, n), make([]int, n)
	copy(a, stack)
	copy(b, stack)

	//find max
	x := -1
	idx := -1
	for i := 0; i < n; i++ {
		if x == -1 {
			if a[i] != 0 && a[i] != 1 {
				x = a[i]
				idx = i
				if i > 0 {
					a[i] = 0
				} else {
					a[i] = 1
				}
			}
		} else {
			if a[i] == x {
				a[i] = a[idx]
			}
		}
	}
	//find min
	y := -1
	for i := 0; i < n; i++ {
		if y == -1 {
			if b[i] != 9 {
				y = b[i]
				b[i] = 9
			}
		} else {
			if b[i] == y {
				b[i] = 9
			}
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res += (b[i] - a[i]) * int(math.Pow(10, float64(n-i-1)))
	}
	return res
}
