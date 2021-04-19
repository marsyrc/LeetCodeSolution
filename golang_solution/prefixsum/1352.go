package prefixsum

type ProductOfNumbers struct {
	arr []int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		arr: []int{1},
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num == 0 {
		this.arr = []int{1}
		return
	}
	last := this.arr[len(this.arr)-1]
	this.arr = append(this.arr, num*last)
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.arr) <= k {
		return 0
	}
	return this.arr[len(this.arr)-1] / this.arr[len(this.arr)-k-1]
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
