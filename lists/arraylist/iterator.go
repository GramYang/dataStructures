package arraylist

type Iterator struct{
	list *List
	index int
}

//返回一个值类型Iterator实例
func (list *List) Iterator() Iterator{
	return Iterator{list:list,index:-1}
}

//累加iterator的index，如果没有越界则返回true
func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	return iterator.list.withinRange(iterator.index)
}

//iterator的index减1，如果没有越界则返回true
func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.list.withinRange(iterator.index)
}

//返回iterator的当前指向的list的值
func(iterator *Iterator) Value() interface{}{
	return iterator.list.elements[iterator.index]
}

//返回iterator当前的index
func (iterator *Iterator) Index() int{
	return iterator.index
}

//重置iterator的index，调用Next()抓取第一个值
func (iterator *Iterator) Begin() {
	iterator.index=-1
}

//置满iterator的index，调用Prev()抓取最后一个值
func (iterator *Iterator)End(){
	iterator.index=iterator.list.size
}

//将iterator指向第一个值
func(iterator *Iterator)First()bool{
	iterator.Begin()
	return iterator.Next()
}

//将iterator指向最后一个值
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}