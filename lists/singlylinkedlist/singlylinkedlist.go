package singlylinkedlist

import (
	"dataStructures/utils"
	"fmt"
	"strings"
)

type List struct {
	first *element
	last  *element
	size  int
}

type element struct {
	value interface{}
	next  *element
}

//新建list，Add进values
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

//添加元素：链接元素，修改last元素，递增size
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.last.next = newElement
			list.last = newElement
		}
		list.size++
	}
}

//和Add一样
func (list *List) Append(values ...interface{}) {
	list.Add(values...)
}

//向前添加元素
func (list *List) Prepend(values ...interface{}) {
	// in reverse to keep passed order i.e. ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &element{value: values[v], next: list.first}
		list.first = newElement
		if list.size == 0 {
			list.last = newElement
		}
		list.size++
	}
}

//根据下标获取元素，需要一个一个的遍历，效率较低
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

//根据下标删除元素，也需要一个一个遍历，并进行链表拼接
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}
	var beforeElement *element
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
		beforeElement = element
	}
	if element == list.first { //删除元素是头元素
		list.first = element.next
	}
	if element == list.last { //删除元素是尾元素
		list.last = beforeElement
	}
	if beforeElement != nil {
		beforeElement.next = element.next
	}
	element = nil
	list.size--
}

//判断list是否包含values，有一个元素不含就返回false，全部都含有才会返回true
func (list *List) Contains(values ...interface{}) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if element.value == value {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

//遍历list，将元素存入一个新建的切片后返回
func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

//返回元素在list中的下标，如果list为空则返回-1，否则遍历list.Values()返回的切片后返回下标
func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for index, element := range list.Values() {
		if element == value {
			return index
		}
	}
	return -1
}

//判断list是否为空
func (list *List) Empty() bool {
	return list.size == 0
}

//返回list的元素数量
func (list *List) Size() int {
	return list.size
}

//清空list
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

//调用list.Values()返回的切片进行排序，再讲切片Add进清空的list
func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	values := list.Values()
	utils.Sort(values, comparator)
	list.Clear()
	list.Add(values...)
}

//交换两个下标节点的值，同样需要遍历list
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) && i != j {
		var element1, element2 *element
		for e, currentElement := 0, list.first; element1 == nil || element2 == nil; e, currentElement = e+1, currentElement.next {
			switch e {
			case i:
				element1 = currentElement
			case j:
				element2 = currentElement
			}
		}
		element1.value, element2.value = element2.value, element1.value
	}
}

//执行下标插入values
func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	list.size += len(values)
	var beforeElement *element
	foundElement := list.first
	for e := 0; e != index; e, foundElement = e+1, foundElement.next {
		beforeElement = foundElement
	}
	if beforeElement == nil {
		return
	}
	if foundElement == list.first {
		oldNextElement := list.first
		for i, value := range values {
			newElement := &element{value: value}
			if i == 0 {
				list.first = newElement
			} else {
				beforeElement.next = newElement //拼接链表
			}
			beforeElement = newElement //顺移下一位
		}
		beforeElement.next = oldNextElement //values的尾部接上index+1元素
	} else {
		oldNextElement := beforeElement.next
		for _, value := range values {
			newElement := &element{value: value}
			beforeElement.next = newElement
			beforeElement = newElement
		}
		beforeElement.next = oldNextElement //values的尾部接上index+1元素
	}
}

//修改制定下标的元素值，同样需要遍历到该下标
func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		// Append
		if index == list.size {
			list.Add(value)
		}
		return
	}
	foundElement := list.first
	for e := 0; e != index; {
		e, foundElement = e+1, foundElement.next
	}
	foundElement.value = value
}

//将list输出成字符串
func (list *List) String() string {
	str := "SinglyLinkedList\n"
	values := []string{}
	for element := list.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}

//判断index是否越界
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}