package util

import (
	"sort"
)

/*
	用于排序的工具包
*/

type Item struct {
	Key   int
	Value int
}

type Items []Item //按照value值排序

func (this Items) Len() int {
	return len(this)
}

//如果value相等,则索引小的数字会排在后面, 表示优先达到的数据更大
func (this Items) Less(i, j int) bool {
	if i < j { //使用Reverse反转后 则索引小的数字会排在前面,表示优先达到的数据更大
		return this[i].Value < this[j].Value
	}
	return this[i].Value <= this[j].Value
}

func (this Items) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this Items) Sort() {
	Sort(this)
}

func (this Items) KSort() {
	KSort(this)
}

//按照value值升序 如果vlaue相等则索引小的在后
func Sort(items Items) {
	sort.Sort(items)
}

//按照value值降序 如果vlaue相等则索引小的在前
func KSort(items Items) {
	a := sort.Reverse(items)
	sort.Sort(a)
}
