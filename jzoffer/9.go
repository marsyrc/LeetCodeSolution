package jianzhioffer

type CQueue struct {
	s1 []int
	s2 []int
}

func Constructor() CQueue {
	s1, s2 := []int{}, []int{}
	return CQueue{
		s1: s1,
		s2: s2,
	}
}

func (this *CQueue) AppendTail(value int) {
	for len(this.s2) != 0 {
		top := this.s2[len(this.s2)-1]
		this.s2 = this.s2[:len(this.s2)-1]
		this.s1 = append(this.s1, top)
	}
	this.s1 = append(this.s1, value)
}

func (this *CQueue) DeleteHead() int {
	for len(this.s1) != 0 {
		top := this.s1[len(this.s1)-1]
		this.s1 = this.s1[:len(this.s1)-1]
		this.s2 = append(this.s2, top)
	}
	if len(this.s2) == 0 {
		return -1
	}
	head := this.s2[len(this.s2)-1]
	this.s2 = this.s2[:len(this.s2)-1]
	return head
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
