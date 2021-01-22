package binarysearch

func findClosestElements(arr []int, k int, x int) []int {
	n := len(arr)
	l, h := 0, n-k
	for l < h {
		mid := l + (h-l)/2
		if x-arr[mid] > arr[mid+k]-x {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return arr[l : l+k]
}
