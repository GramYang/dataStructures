package doublylinkedlist

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
	prev  *element
	next  *element
}

//新建一个list
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

//向list后拼接values
func (list *List) Add(values ...interface{}) {
	for _, value := range values {
		newElement := &element{value: value, prev: list.last}
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

//就是Add
func (list *List) Append(values ...interface{}) {
	list.Add(values...)
}

//向list前拼接values
func (list *List) Prepend(values ...interface{}) {
	for v := len(values) - 1; v >= 0; v-- {
		newElement := &element{value: values[v], next: list.first}
		if list.size == 0 {
			list.first = newElement
			list.last = newElement
		} else {
			list.first.prev = newElement
			list.first = newElement
		}
		list.size++
	}
}

//根据下标返回元素，成功则返回元素和true，失败则返回nil和false
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	if list.size-index < index {
		element := list.last
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
		}
		return element.value, true
	}
	element := list.first
	for e := 0; e != index; e, element = e+1, element.next {
	}
	return element.value, true
}

//删除list指定下标元素，从头或尾查找节点，然后移除当前节点
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	if list.size == 1 {
		list.Clear()
		return
	}
	var element *element
	if list.size-index < index {
		element = list.last
		for e := list.size - 1; e != index; e, element = e-1, element.prev {
		}
	} else {
		element = list.first
		for e := 0; e != index; e, element = e+1, element.next {
		}
	}
	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = element.prev
	}
	if element.prev != nil {
		element.prev.next = element.next
	}
	if element.next != nil {
		element.next.prev = element.prev
	}
	element = nil
	list.size--
}

//检查list含有values中的一项，返回true
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

//遍历list将所有元素添加进新建的切片
func (list *List) Values() []interface{} {
	values := make([]interface{}, list.size, list.size)
	for e, element := 0, list.first; element != nil; e, element = e+1, element.next {
		values[e] = element.value
	}
	return values
}

//本来list是没有index的，这里返回的是list.Values()返回的切片的index
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

//返回list尺寸
func (list *List) Size() int {
	return list.size
}

//清空list
func (list *List) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

//将list转换为切片，然后排序，排序完后清空list再重新添加
func (list *List) Sort(comparator utils.Comparator) {
	if list.size < 2 {
		return
	}
	values := list.Values()
	utils.Sort(values, comparator)
	list.Clear()
	list.Add(values...)
}

//先从头遍历到i和j位置找到两个节点，然后交换他们的值
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

//向index位置插入values。先遍历发现index的节点，然后遍历values进行添加
func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(values...)
		}
		return
	}
	list.size += len(values)
	var beforeElement *element
	var foundElement *element
	if list.size-index < index {
		foundElement = list.last
		for e := list.size - 1; e != index; e, foundElement = e-1, foundElement.prev {
			beforeElement = foundElement.prev
		}
	} else {
		foundElement = list.first
		for e := 0; e != index; e, foundElement = e+1, foundElement.next {
			beforeElement = foundElement
		}
	}
	if beforeElement==nil {
		return
	}
	if foundElement == list.first {
		oldNextElement := list.first
		for i, value := range values {
			newElement := &element{value: value}
			if i == 0 {
				list.first = newElement
			} else {
				newElement.prev = beforeElement
				beforeElement.next = newElement
			}
			beforeElement = newElement
		}
		oldNextElement.prev = beforeElement
		beforeElement.next = oldNextElement
	} else {
		oldNextElement := beforeElement.next
		for _, value := range values {
			newElement := &element{value: value}
			newElement.prev = beforeElement
			beforeElement.next = newElement
			beforeElement = newElement
		}
		oldNextElement.prev = beforeElement
		beforeElement.next = oldNextElement
	}
}

//修改index位置节点的值为value，也是遍历发现index节点
func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		if index == list.size {
			list.Add(value)
		}
		return
	}
	var foundElement *element
	if list.size-index < index {
		foundElement = list.last
		for e := list.size - 1; e != index; {
			//fmt.Println("Set last", index, value, foundElement, foundElement.prev)
			e, foundElement = e-1, foundElement.prev
		}
	} else {
		foundElement = list.first
		for e := 0; e != index; {
			e, foundElement = e+1, foundElement.next
		}
	}
	foundElement.value = value
}

//将list输出成字符串
func (list *List) String() string {
	str := "DoublyLinkedList\n"
	values := []string{}
	for element := list.first; element != nil; element = element.next {
		values = append(values, fmt.Sprintf("%v", element.value))
	}
	str += strings.Join(values, ", ")
	return str
}

//index没有越界返回true
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}