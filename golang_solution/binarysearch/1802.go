package binarysearch

func maxValue(n int, index int, maxSum int) int {
	l, r := 0, 1000000001
	ans := 0
	for l < r {
		mid := (l + r) / 2
		if value(n, index, maxSum, mid) {
			ans = mid
			l = mid + 1
		} else {
			r = mid
		}
	}
	return ans
}

func value(n int, index int, maxSum int, target int) bool {
	sum := 0
	k := 0
	if target > index+1 {
		sum += (target - index + target) * (index + 1) / 2
		k += index + 1
	} else {
		sum += (target + 1) * target / 2
		k += target
	}
	if n-index >= target {
		sum += (target + 1) * target / 2
		k += target
	} else {
		sum += (target - n + index + target + 1) * (n - index) / 2
		k += n - index
	}
	sum = sum - target + (n - k + 1)
	return maxSum >= sum
}
