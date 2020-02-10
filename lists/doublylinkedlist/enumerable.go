package doublylinkedlist

//遍历iterator，执行f(k,v)
func (list *List) Each(f func(index int, value interface{})) {
	iterator := list.Iterator()
	for iterator.Next() {
		f(iterator.Index(), iterator.Value())
	}
}

//遍历iterator，执行f(k,v)后返回的结果Add进一个新建的list并返回
func (list *List) Map(f func(index int, value interface{}) interface{}) *List {
	newList := &List{}
	iterator := list.Iterator()
	for iterator.Next() {
		newList.Add(f(iterator.Index(), iterator.Value()))
	}
	return newList
}

//遍历iterator，执行f(k,v)返回结果为true则将其值Add进新建的list并返回
func (list *List) Select(f func(index int, value interface{}) bool) *List {
	newList := &List{}
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			newList.Add(iterator.Value())
		}
	}
	return newList
}

//遍历iterator，执行f(k,v)遇到一个返回true则返回true，一个没遇到就返回false
func (list *List) Any(f func(index int, value interface{}) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return true
		}
	}
	return false
}

//和上面相反
func (list *List) All(f func(index int, value interface{}) bool) bool {
	iterator := list.Iterator()
	for iterator.Next() {
		if !f(iterator.Index(), iterator.Value()) {
			return false
		}
	}
	return true
}

//遍历iterator执行f(k,v)，如果有返回true则返回iterator的index和value，没有遇到true则返回-1和nil
func (list *List) Find(f func(index int, value interface{}) bool) (index int, value interface{}) {
	iterator := list.Iterator()
	for iterator.Next() {
		if f(iterator.Index(), iterator.Value()) {
			return iterator.Index(), iterator.Value()
		}
	}
	return -1, nil
}