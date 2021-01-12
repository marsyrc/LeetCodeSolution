package main

import "fmt"

// 对所有待排序元素建堆，然后依次取出堆顶元素，就可以得到排好序的序列。
// 当当前的结点下标为 i 时，父结点、左子结点和右子结点的选择方式如下：

// 这里 floor 函数将实数映射到最小的前导整数。
// iParent(i) = floor((i - 1) / 2);
// iLeftChild(i) = 2 * i + 1;
// iRightChild(i) = 2 * i + 2;

//runtime : nlogn
func bulidHeap(arr []int, start int, end int) {
	dad := start
	son := dad*2 + 1
	for son <= end {
		if son+1 <= end && arr[son] < arr[son+1] {
			son++
		}
		if arr[dad] > arr[son] {
			return
		} else {
			swap(arr, dad, son)
			dad = son
			son = dad*2 + 1
		}
	}
}

func heapSort(arr []int, len int) {
	for i := len/2 - 1; i >= 0; i-- {
		bulidHeap(arr, i, len-1)
	}
	for i := len - 1; i > 0; i-- {
		swap(arr, 0, i)
		bulidHeap(arr, 0, i-1)
	}
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func main() {
	a := []int{7, 5, 3, 7, 3, 26, 2, 1, 23}
	heapSort(a, len(a))
	fmt.Println(a)
}
