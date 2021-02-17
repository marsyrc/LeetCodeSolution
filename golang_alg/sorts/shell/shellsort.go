package main

import "fmt"

func shellSort(a []int) {
	//生成步长
	step := 0
	for step <= len(a) {
		step = 3*step + 1
	}

	for step > 0 {
		for i := step; i < len(a); i++ {
			j := i - step
			backup := a[i]
			for j >= 0 && a[j] > backup {
				a[j+step] = a[j]
				j = j - step
			}
			a[j+step] = backup
		}
		fmt.Println(step)
		step = (step - 1) / 3
	}
}

func main() {
	a := []int{1, 2, 5, 4, 2, 9}
	shellSort(a)
	fmt.Println(a)
}
