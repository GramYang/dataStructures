package avltree

import (
	"dataStructures/utils"
	"encoding/json"
)

//将树输出为json键值对，其中以元素key值小到大排列
func (tree *Tree) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	it := tree.Iterator()
	for it.Next() {
		elements[utils.ToString(it.Key())] = it.Value()
	}
	return json.Marshal(&elements)
}

//将json键值对转换为树，重新构建avl树
func (tree *Tree) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		tree.Clear()
		for key, value := range elements {
			tree.Put(key, value)
		}
	}
	return err
}