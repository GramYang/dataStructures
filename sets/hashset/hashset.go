package hashset

import (
	"fmt"
	"strings"
)

// Set holds elements in go's native map
type Set struct {
	items map[interface{}]struct{}
}

var itemExists = struct{}{}

//返回set实例，用values初始化
func New(values ...interface{}) *Set {
	set := &Set{items: make(map[interface{}]struct{})}
	if len(values) > 0 {
		set.Add(values...)
	}
	return set
}

//添加items，其中key是item，value是空结构体
func (set *Set) Add(items ...interface{}) {
	for _, item := range items {
		set.items[item] = itemExists
	}
}

//删除items，删除key为item的键值对
func (set *Set) Remove(items ...interface{}) {
	for _, item := range items {
		delete(set.items, item)
	}
}

//判断set是否包含items
func (set *Set) Contains(items ...interface{}) bool {
	for _, item := range items {
		if _, contains := set.items[item]; !contains {
			return false
		}
	}
	return true
}

//判断set是否为空
func (set *Set) Empty() bool {
	return set.Size() == 0
}

//返回set大小
func (set *Set) Size() int {
	return len(set.items)
}

//清空set
func (set *Set) Clear() {
	set.items = make(map[interface{}]struct{})
}

//将set中的key存入切片后返回
func (set *Set) Values() []interface{} {
	values := make([]interface{}, set.Size())
	count := 0
	for item := range set.items {
		values[count] = item
		count++
	}
	return values
}

//将set的key值输出为字符串
func (set *Set) String() string {
	str := "HashSet\n"
	items := []string{}
	for k := range set.items {
		items = append(items, fmt.Sprintf("%v", k))
	}
	str += strings.Join(items, ", ")
	return str
}