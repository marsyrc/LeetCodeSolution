package jianzhioffer

func getLeastNumbers(arr []int, k int) []int {
	if k == 0 { //k=0的情况不能忘记
		return []int{}
	}
	start, end := 0, len(arr)-1
	index := partition(arr, start, end)
	for index != k-1 { //一旦index==k-1，即可停止。
		if index > k-1 {
			end = index - 1
		} else {
			start = index + 1
		}
		index = partition(arr, start, end)
	}
	return arr[:k] //注意：返回的这k个元素不一定是有序的
}

func partition(arr []int, left, right int) int {
	i, j := left, right
	pivot := arr[left]
	for {
		for arr[i] <= pivot && i < right {
			i++
		}
		for arr[j] >= pivot && j > left {
			j--
		}
		if i < j {
			swap(arr, i, j)
		} else {
			break
		}
	}
	swap(arr, left, j)
	return j
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
