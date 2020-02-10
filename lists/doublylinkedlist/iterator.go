package doublylinkedlist

type Iterator struct {
	list    *List
	index   int
	element *element
}

//生成Iterator值实例
func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1, element: nil}
}

//iterator指向下一个元素，递增index，返回true则element指向下一个元素，返回false则element指向nil
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index != 0 {
		iterator.element = iterator.element.next
	} else {
		iterator.element = iterator.list.first
	}
	return true
}

//iterator指向前一个元素，递减index，返回true则
func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	if !iterator.list.withinRange(iterator.index) {
		iterator.element = nil
		return false
	}
	if iterator.index == iterator.list.size-1 {
		iterator.element = iterator.list.last
	} else {
		iterator.element = iterator.element.prev
	}
	return iterator.list.withinRange(iterator.index)
}

//返回iterator指向element的值
func (iterator *Iterator) Value() interface{} {
	return iterator.element.value
}

//返回iterator的index
func (iterator *Iterator) Index() int {
	return iterator.index
}

//将iterator设置到初始状态
func (iterator *Iterator) Begin() {
	iterator.index = -1
	iterator.element = nil
}

//将iterator设置到末尾
func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
	iterator.element = iterator.list.last
}

//iterator移向第一个元素，如果有元素则返回true
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

//返回true说明iterator指向倒数第二个元素
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}