package circlebuffer

func (buffer *Buffer) Each(f func(index int, value interface{})){
	iterator:=buffer.Iterator()
	for iterator.Next(){
		f(iterator.Index(),iterator.Value())
	}
}

func(buffer *Buffer) Map(f func(index int, value interface{}) interface{}) *Buffer{
	newBuffer:=&Buffer{}
	iterator:=buffer.Iterator()
	for iterator.Next(){
		newBuffer.Add(f(iterator.Index(),iterator.Value()))
	}
	return newBuffer
}

func(buffer *Buffer)Select(f func(index int, value interface{}) bool) *Buffer{
	newBuffer:=&Buffer{}
	iterator:=buffer.Iterator()
	for iterator.Next(){
		if f(iterator.Index(), iterator.Value()){
			newBuffer.Add(iterator.Value())
		}
	}
	return newBuffer
}

func (buffer *Buffer)Any(f func(index int,value interface{})bool) bool{
	iterator:=buffer.Iterator()
	for iterator.Next(){
		if f(iterator.Index(),iterator.Value()){
			return true
		}
	}
	return false
}

func(buffer *Buffer) All(f func(index int, value interface{})bool)bool{
	iterator:=buffer.Iterator()
	for iterator.Next(){
		if !f(iterator.Index(),iterator.Value()){
			return false
		}
	}
	return true
}

func(buffer *Buffer)Find(f func(index int,value interface{})bool) (int, interface{}){
	iterator:=buffer.Iterator()
	for iterator.Next(){
		if f(iterator.Index(),iterator.Value()){
			return iterator.Index(),iterator.Value()
		}
	}
	return -1,nil
}