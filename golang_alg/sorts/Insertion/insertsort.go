package Insertion

/*
   in place way
*/

func Sort(arr []int) []int {
	n := len(arr)
	if n == 1 {
		return arr
	}
	for i := 1; i < n; i++ {
		backup := arr[i]
		j := i - 1
		for j >= 0 && backup < arr[j] {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = backup
	}
	return arr
}
