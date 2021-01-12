package main

import (
	"fmt"
	"math"
)

func insertSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		backup := arr[i]
		j := i - 1
		for j >= 0 && arr[j] >= backup {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = backup
	}
}

func bucketSort(arr []int) {
	n := len(arr)
	maxNum := max(arr...)
	buckets := make([][]int, n)

	index := 0
	for i := 0; i < n; i++ {
		index = arr[i] * (n - 1) / maxNum
		buckets[index] = append(buckets[index], arr[i])
	}

	tempPos := 0
	for i := 0; i < n; i++ {
		if len(buckets[i]) > 0 {
			insertSort(buckets[i])
			copy(arr[tempPos:], buckets[i])
			tempPos += len(buckets[i])
		}
	}
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

func main() {
	a := []int{7, 5, 3, 7, 3, 26, 2, 1, 23}
	bucketSort(a)
	fmt.Println(a)
}
