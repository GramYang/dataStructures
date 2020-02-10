package circlebuffer

type Iterator struct {
	buffer *Buffer
	index int
}

func (buffer *Buffer) Iterator() Iterator{
	return Iterator{buffer:buffer,index:-1}
}

func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.buffer.Length {
		iterator.index++
	}
	return iterator.buffer.withinRange(iterator.index)
}

func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.buffer.withinRange(iterator.index)
}

func(iterator *Iterator) Value() interface{}{
	return iterator.buffer.elements[iterator.index]
}

func (iterator *Iterator) Index() int{
	return iterator.index
}

func (iterator *Iterator) Begin() {
	iterator.index=-1
}

func (iterator *Iterator)End(){
	iterator.index=iterator.buffer.Length
}

func(iterator *Iterator)First()bool{
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}