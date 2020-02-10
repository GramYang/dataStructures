package binaryheap

type Iterator struct {
	heap  *Heap
	index int
}

//返回迭代器实例
func (heap *Heap) Iterator() Iterator {
	return Iterator{heap: heap, index: -1}
}

//递增迭代器的index，返回true表示没有越界
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.heap.Size() {
		iterator.index++
	}
	return iterator.heap.withinRange(iterator.index)
}

//递减迭代器的index，返回true表示没有越界
func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.heap.withinRange(iterator.index)
}

//获取迭代器指向的值
func (iterator *Iterator) Value() interface{} {
	value, _ := iterator.heap.list.Get(iterator.index)
	return value
}

//返回迭代器的下标
func (iterator *Iterator) Index() int {
	return iterator.index
}

//初始化迭代器下标
func (iterator *Iterator) Begin() {
	iterator.index = -1
}

//迭代器下标指向末尾后一位
func (iterator *Iterator) End() {
	iterator.index = iterator.heap.Size()
}

//迭代器指向第一个元素
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

//迭代器指向最后一个元素
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}