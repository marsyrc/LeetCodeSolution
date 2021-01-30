package design

import (
	"math/rand"
	"time"
)

type RandomizedCollection struct {
	m    map[int]map[int]struct{}
	nums []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedCollection {
	return RandomizedCollection{
		m: make(map[int]map[int]struct{}),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.nums = append(this.nums, val)
	_, ok := this.m[val]
	if !ok {
		this.m[val] = map[int]struct{}{}
		this.m[val][len(this.nums)-1] = struct{}{}
	} else {
		this.m[val][len(this.nums)-1] = struct{}{}
	}
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	_, ok := this.m[val]
	if !ok {
		return false
	}
	curIndex := -1
	for i := range this.m[val] {
		curIndex = i
		break
	}
	lastIndex := len(this.nums) - 1
	lastVal := this.nums[lastIndex]

	//用最后一个替换
	this.nums[curIndex] = lastVal
	delete(this.m[val], curIndex)

	//修正索引
	if _, ok := this.m[lastVal]; !ok {
		this.m[lastVal] = map[int]struct{}{}
	}
	if curIndex < lastIndex {
		this.m[lastVal][curIndex] = struct{}{}
	}

	if len(this.m[val]) == 0 {
		delete(this.m, val)
	}
	delete(this.m[lastVal], lastIndex)
	this.nums = this.nums[:lastIndex]
	return true
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	rand.Seed(time.Now().UnixNano())
	return this.nums[rand.Intn(len(this.nums))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
