package jianzhioffer

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1
	for l < r {
		mid := l + (r-l)/2
		if numbers[mid] < numbers[r] {
			r = mid
		} else if numbers[mid] > numbers[r] {
			l = mid + 1
		} else {
			r--
		}
	}
	return numbers[l]
}
