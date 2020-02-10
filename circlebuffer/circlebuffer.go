package circlebuffer

import (
	"dataStructures/utils"
	"fmt"
	"strings"
)

type Buffer struct {
	elements []interface{}
	Length     int //隐藏elements
	isFull bool
	pointer  int
}

func New(len int) *Buffer{
	buffer := &Buffer{elements:make([]interface{},len,len), Length:len,isFull:false,pointer:0}
	return buffer
}

func (buffer *Buffer) Add(values ...interface{}) {
	if !buffer.isFull && buffer.pointer+len(values)>=buffer.Length{
		buffer.isFull=true
	}
	if len(values) + buffer.pointer <= buffer.Length {
		for _,v :=range values {
			buffer.elements[buffer.pointer]=v
			buffer.pointer++
		}
	} else {
		reallen:=(len(values)-(buffer.Length -buffer.pointer))%buffer.Length
		newslice:=values[0:reallen]
		buffer.pointer=0
		for _,v:=range newslice {
			buffer.elements[buffer.pointer]=v
			buffer.pointer++
		}
	}
}

func (buffer *Buffer) Get(index int) (interface{},bool) {
	if!buffer.withinRange(index){
		return nil,false
	}
	return buffer.elements[index],true
}

func (buffer *Buffer) Contains(values ...interface{}) bool {
	for _,searchValue :=range values {
		found:=false
		for _,element:=range buffer.elements{
			if element==searchValue{
				found=true
				break
			}
		}
		if !found{
			return false
		}
	}
	return true
}

func (buffer *Buffer) Values() []interface{} {
	newElements:=make([]interface{},buffer.Length,buffer.Length)
	copy(newElements,buffer.elements[:buffer.Length])
	return newElements
}

func (buffer *Buffer) IndexOf(value interface{}) int{
	if buffer.Length ==0 {
		return -1
	}
	for index, element := range buffer.elements{
		if element==value {
			return index
		}
	}
	return -1
}

func (buffer *Buffer) Empty() bool {
	return buffer.Length ==0
}

func (buffer *Buffer) Size() int{
	if buffer.isFull {
		return len(buffer.elements)
	}
	return buffer.pointer
}

func (buffer *Buffer) Clear() {
	buffer.Length =0
	buffer.elements=buffer.elements[0:0]
}

func (buffer *Buffer) Sort(comparator utils.Comparator) {
	if len(buffer.elements)<2{
		return
	}
	utils.Sort(buffer.elements[:buffer.Length],comparator)
}

func (buffer *Buffer) Swap(i, j int) {
	if buffer.withinRange(i)&&buffer.withinRange(j){
		buffer.elements[i],buffer.elements[j]=buffer.elements[j],buffer.elements[i]
	}
}

func (buffer *Buffer) Set(index int, value interface{}){
	if !buffer.withinRange(index){
		return
	}
	buffer.elements[index] = value
}

func (buffer *Buffer) String() string{
	str := "CircleBuffer\n"
	values := []string{}
	for _, value := range buffer.elements[:buffer.Length] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (buffer *Buffer) withinRange(index int) bool{
	return index>=0&&index<buffer.Length
}