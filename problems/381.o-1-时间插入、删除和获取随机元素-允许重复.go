package problems

import "math/rand"

/*
 * @lc app=leetcode.cn id=381 lang=golang
 *
 * [381] O(1) 时间插入、删除和获取随机元素 - 允许重复
 */

// @lc code=start
type RandomizedCollection struct {
	items []int
	m     map[int]map[int]struct{}
}

/** Initialize your data structure here. */
func Constructor_() RandomizedCollection {
	return RandomizedCollection{
		m: make(map[int]map[int]struct{}),
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	idx, exists := this.m[val]
	if !exists {
		idx = make(map[int]struct{})
	}

	this.items = append(this.items, val)
	idx[len(this.items)-1] = struct{}{}
	this.m[val] = idx
	return !exists

}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	idx, exists := this.m[val]
	if !exists {
		return false
	}
	var i int
	for id := range idx {
		i = id
		break
	}
	n := len(this.items)
	this.items[i] = this.items[n-1]
	delete(idx, i)
	delete(this.m[this.items[i]], n-1)
	if i < n-1 {
		this.m[this.items[i]][i] = struct{}{}
	}
	if len(idx) == 0 {
		delete(this.m, val)
	}
	this.items = this.items[:n-1]
	return true

}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.items[rand.Intn(len(this.items))]
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
