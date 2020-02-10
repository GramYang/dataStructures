package binaryheap

import (
	"dataStructures/lists/arraylist"
	"dataStructures/utils"
	"fmt"
	"strings"
)

type Heap struct {
	list       *arraylist.List
	Comparator utils.Comparator
}

//新建一个堆返回
func NewWith(comparator utils.Comparator) *Heap {
	return &Heap{list: arraylist.New(), Comparator: comparator}
}

//新建一个堆返回，堆内元素为int
func NewWithIntComparator() *Heap {
	return &Heap{list: arraylist.New(), Comparator: utils.IntComparator}
}

//新建一个堆返回，堆内元素为string
func NewWithStringComparator() *Heap {
	return &Heap{list: arraylist.New(), Comparator: utils.StringComparator}
}

//存入values，然后bubble up
func (heap *Heap) Push(values ...interface{}) {
	if len(values) == 1 {
		heap.list.Add(values[0])
		heap.bubbleUp()
	} else {
		// Reference: https://en.wikipedia.org/wiki/Binary_heap#Building_a_heap
		for _, value := range values {
			heap.list.Add(value)
		}
		size := heap.list.Size()/2 + 1
		for i := size; i >= 0; i-- {
			heap.bubbleDownIndex(i)
		}
	}
}

//删除堆顶元素并返回
func (heap *Heap) Pop() (value interface{}, ok bool) {
	value, ok = heap.list.Get(0)
	if !ok {
		return
	}
	lastIndex := heap.list.Size() - 1
	heap.list.Swap(0, lastIndex)
	heap.list.Remove(lastIndex)
	heap.bubbleDown()
	return
}

//返回堆顶元素且不删除
func (heap *Heap) Peek() (value interface{}, ok bool) {
	return heap.list.Get(0)
}

//堆为空则返回true
func (heap *Heap) Empty() bool {
	return heap.list.Empty()
}

//返回堆容积
func (heap *Heap) Size() int {
	return heap.list.Size()
}

//清空堆
func (heap *Heap) Clear() {
	heap.list.Clear()
}

//将堆内值作为切片返回
func (heap *Heap) Values() []interface{} {
	return heap.list.Values()
}

//将堆输出为字符串
func (heap *Heap) String() string {
	str := "BinaryHeap\n"
	values := []string{}
	for _, value := range heap.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Performs the "bubble down" operation. This is to place the element that is at the root
// of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *Heap) bubbleDown() {
	heap.bubbleDownIndex(0)
}

// Performs the "bubble down" operation. This is to place the element that is at the index
// of the heap in its correct place so that the heap maintains the min/max-heap order property.
func (heap *Heap) bubbleDownIndex(index int) {
	size := heap.list.Size()
	for leftIndex := index<<1 + 1; leftIndex < size; leftIndex = index<<1 + 1 {
		rightIndex := index<<1 + 2
		smallerIndex := leftIndex
		leftValue, _ := heap.list.Get(leftIndex)
		rightValue, _ := heap.list.Get(rightIndex)
		if rightIndex < size && heap.Comparator(leftValue, rightValue) > 0 {
			smallerIndex = rightIndex
		}
		indexValue, _ := heap.list.Get(index)
		smallerValue, _ := heap.list.Get(smallerIndex)
		if heap.Comparator(indexValue, smallerValue) > 0 {
			heap.list.Swap(index, smallerIndex)
		} else {
			break
		}
		index = smallerIndex
	}
}

// Performs the "bubble up" operation. This is to place a newly inserted
// element (i.e. last element in the list) in its correct place so that
// the heap maintains the min/max-heap order property.
func (heap *Heap) bubbleUp() {
	index := heap.list.Size() - 1
	for parentIndex := (index - 1) >> 1; index > 0; parentIndex = (index - 1) >> 1 {
		indexValue, _ := heap.list.Get(index)
		parentValue, _ := heap.list.Get(parentIndex)
		if heap.Comparator(parentValue, indexValue) <= 0 {
			break
		}
		heap.list.Swap(index, parentIndex)
		index = parentIndex
	}
}

//检查index是否越界
func (heap *Heap) withinRange(index int) bool {
	return index >= 0 && index < heap.list.Size()
}