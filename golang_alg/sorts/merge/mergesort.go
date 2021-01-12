package merge

import "fmt"

//分治 稳定 非原地
//runtime nlogn
func mergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	middle := length / 2
	left := arr[:middle]
	right := arr[middle:]
	return merge(mergeSort(left), mergeSort(right))
}

func merge(left []int, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}

	if len(left) > 0 {
		result = append(result, left...)
	}
	if len(right) > 0 {
		result = append(result, right...)
	}
	return result
}

func main() {
	a := []int{7, 5, 3, 7, 3, 26, 2, 1, 23}
	fmt.Println(mergeSort(a))
}
