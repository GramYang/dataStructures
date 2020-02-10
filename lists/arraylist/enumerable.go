package arraylist

//遍历list，调用f(k,v)
func (list *List) Each(f func(index int, value interface{})){
	iterator:=list.Iterator()
	for iterator.Next(){
		f(iterator.Index(),iterator.Value())
	}
}

//遍历list，调用f(k,v)后返回的值存入一个新的list并返回
func(list *List) Map(f func(index int, value interface{}) interface{}) *List{
	newList:=&List{}
	iterator:=list.Iterator()
	for iterator.Next(){
		newList.Add(f(iterator.Index(),iterator.Value()))
	}
	return newList
}

//遍历list，执行f(k,v)返回结果为true则将v存入一个新list并返回
func(list *List)Select(f func(index int, value interface{}) bool) *List{
	newList:=&List{}
	iterator:=list.Iterator()
	for iterator.Next(){
		if f(iterator.Index(), iterator.Value()){
			newList.Add(iterator.Value())
		}
	}
	return newList
}

//遍历list，执行f(k,v)有返回true则返回true
func (list *List)Any(f func(index int,value interface{})bool) bool{
	iterator:=list.Iterator()
	for iterator.Next(){
		if f(iterator.Index(),iterator.Value()){
			return true
		}
	}
	return false
}

//遍历list，执行f(k,v)有返回false则返回false
func(list *List) All(f func(index int, value interface{})bool)bool{
	iterator:=list.Iterator()
	for iterator.Next(){
		if !f(iterator.Index(),iterator.Value()){
			return false
		}
	}
	return true
}

//遍历list，执行f(k,v)有返回true则返回k和v
func(list *List)Find(f func(index int,value interface{})bool) (int, interface{}){
	iterator:=list.Iterator()
	for iterator.Next(){
		if f(iterator.Index(),iterator.Value()){
			return iterator.Index(),iterator.Value()
		}
	}
	return -1,nil
}