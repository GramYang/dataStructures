package singlylinkedlist


type Iterator struct {
	list    *List
	index   int
	element *element
}

//返回一个迭代器值实例
func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1, element: nil}
}

//迭代器递增，如果越界了返回false，否则迭代器指向某一元素且返回true
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index == 0 {
		iterator.element = iterator.list.first
	} else {
		iterator.element = iterator.element.next
	}
	return true
}

//返回迭代器指向元素的值
func (iterator *Iterator) Value() interface{} {
	return iterator.element.value
}

//返回迭代器当前下标
func (iterator *Iterator) Index() int {
	return iterator.index
}

//迭代器回归初始状态
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

//迭代器指向第一个元素
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}