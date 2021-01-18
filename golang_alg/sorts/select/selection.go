package main

import "fmt"

func selectSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}

func main() {
	a := []int{1, 2, 42, 14, 12, 4, 2}
	selectSort(a)
	fmt.Println(a)
}
