package dynamicplanning

//use prefix sum array to record sum of [0, i]
type NumArray struct {
	prefix []int
	nums   []int
}

func Constructor(nums []int) NumArray {
	if len(nums) == 0 {
		return NumArray{}
	}
	pre := make([]int, len(nums))
	pre[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		pre[i] = pre[i-1] + nums[i]
	}
	newNumArray := NumArray{
		prefix: pre,
		nums:   nums,
	}
	return newNumArray
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.prefix[j] - this.prefix[i] + this.nums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
