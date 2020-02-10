package arraylist

import (
	"dataStructures/utils"
	"fmt"
	"strings"
)

const (
	growthFactor=float32(2.0)//扩容因数
	shrinkFactor=float32(0.25)//收缩因数
)

type List struct {
	elements []interface{}
	size int
}

func New(values ...interface{}) *List {
	list:=&List{}
	if len(values)>0 {
		list.Add(values...)
	}
	return list
}

//添加一个可变参数列表values，直接将values拼接到list.elements
func (list *List) Add(values ...interface{}) {
	list.growBy(len(values))
	for _,value:=range values{
		list.elements[list.size]=value
		list.size++
	}
}

//根据下标返回元素，并检查下标是否越界
func (list *List) Get(index int) (interface{},bool) {
	if!list.withinRange(index){
		return nil,false
	}
	return list.elements[index],true
}

//根据下标index删除元素，检查index是否越界然后将index右侧的切片整体左移一位
//关于耗时操作，Java里面也是用的System.arraycopy()，一样的做法。
func (list *List) Remove(index int) {
	if !list.withinRange(index){
		return
	}
	list.elements[index]=nil
	copy(list.elements[index:],list.elements[index+1:list.size]) //左移一位，耗时操作，需要优化
	list.size--
	list.shrink()
}

//检查list中是否包含values中的任意一个
func (list *List) Contains(values ...interface{}) bool {
	for _,searchValue:=range values{
		found:=false
		for _,element:=range list.elements{
			if element==searchValue{
				found=true
				break
			}
		}
		if !found{
			return false
		}
	}
	return true //values为空也返回true
}

//返回一个切片，里面复制了list.elements
func (list *List) Values() []interface{} {
	newElements:=make([]interface{},list.size,list.size)
	copy(newElements,list.elements[:list.size])
	return newElements
}

//返回元素的下标，list为空或者不包含value都返回-1
func (list *List) IndexOf(value interface{}) int{
	if list.size==0{
		return -1
	}
	for index,element:=range list.elements{
		if element==value{
			return index
		}
	}
	return -1
}

//判断list是否为空
func (list *List) Empty() bool {
	return list.size==0
}

//返回list的大小
func (list *List) Size() int{
	return list.size
}

//清空list
func (list *List) Clear() {
	list.size=0
	list.elements=[]interface{}{}
}

//list排序
func (list *List) Sort(comparator utils.Comparator) {
	if len(list.elements)<2{
		return
	}
	utils.Sort(list.elements[:list.size],comparator)
}

//交换两个指定下标的元素
func (list *List) Swap(i, j int) {
	if list.withinRange(i)&&list.withinRange(j){
		list.elements[i],list.elements[j]=list.elements[j],list.elements[i]
	}
}

//指定下标index插入values，index右元素要右移
func (list *List) Insert(index int, values ...interface{}){
	if !list.withinRange(index){
		if index==list.size{
			list.Add(values)
		}
		return
	}
	l:=len(values)
	list.growBy(l)
	list.size+=l
	copy(list.elements[index+l:],list.elements[index:list.size-l])
	copy(list.elements[index:],values)
}

//修改index下标的值
func (list *List) Set(index int, value interface{}){
	if !list.withinRange(index){
		if index==list.size{
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

//将list按照格式转化为string
func (list *List) String() string{
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

//检查是否需要扩容，需要则根据扩容因子扩容
func (list *List) growBy(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n>=currentCapacity{
		newCapacity:=int(growthFactor*float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

//新建一个cap长度的空切片，将原来的elements复制进去
func (list *List) resize(cap int) {
	newElements:=make([]interface{},cap,cap)
	copy(newElements,list.elements)
	list.elements=newElements
}

//检查index是否越界
func (list *List) withinRange(index int) bool {
	return index>=0&&index<list.size
}

//收缩容量
func (list *List) shrink() {
	if shrinkFactor==0.0{
		return
	}
	currentCapacity:=cap(list.elements)
	if list.size<=int(float32(currentCapacity)*shrinkFactor){
		list.resize(list.size)
	}
}