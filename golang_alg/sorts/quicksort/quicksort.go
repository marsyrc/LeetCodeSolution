package main

import "fmt"

func quickSort(arr []int, left, right int) {
	if left < right {
		index := partition(arr, left, right)
		quickSort(arr, left, index-1)
		quickSort(arr, index+1, right)
	}
}

func grayCode(n int) []int {
	res := []int{0}
	for cur := 0; cur < n; cur++ {
		length := len(res)
		for i := length - 1; i >= 0; i-- { //对之前的格雷码进行逆序并高位补1
			res = append(res, res[i]+(1<<cur))
		}
	}
	return res
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

func main() {
	a := []int{7, 5, 3, 7, 3, 26, 2, 1, 23}
	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}
