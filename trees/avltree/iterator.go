package avltree

type Iterator struct {
	tree     *Tree
	node     *Node
	position position
}

type position byte

const (
	begin, between, end position = 0, 1, 2
)

//返回迭代器实例
func (tree *Tree) Iterator() *Iterator {
	return &Iterator{tree: tree, node: nil, position: begin}
}

//遍历avl树，改变状态，返回true说明node不为空
func (iterator *Iterator) Next() bool {
	switch iterator.position {
	case begin: //初始状态在Next()后指向树中最左叶子
		iterator.position = between
		iterator.node = iterator.tree.Left()
	case between: //遍历最近的一个大于当前节点的节点，这个节点可能是子节点也可能是父节点
		iterator.node = iterator.node.Next()
	}
	if iterator.node == nil {//end状态
		iterator.position = end
		return false
	}
	return true
}

//反向遍历avl树，返回true说明node不为空
func (iterator *Iterator) Prev() bool {
	switch iterator.position {
	case end:
		iterator.position = between
		iterator.node = iterator.tree.Right()
	case between:
		iterator.node = iterator.node.Prev()
	}

	if iterator.node == nil {
		iterator.position = begin
		return false
	}
	return true
}

//返回iterator指向节点的value值
func (iterator *Iterator) Value() interface{} {
	if iterator.node == nil {
		return nil
	}
	return iterator.node.Value
}

//返回iterator指向节点的key值
func (iterator *Iterator) Key() interface{} {
	if iterator.node == nil {
		return nil
	}
	return iterator.node.Key
}

//iterator设置到begin
func (iterator *Iterator) Begin() {
	iterator.node = nil
	iterator.position = begin
}

//iterator设置到end
func (iterator *Iterator) End() {
	iterator.node = nil
	iterator.position = end
}

//iterator指向avl树的第一个值，也就是最左叶子
func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

//iterator指向avl树的第后一个值，也就是最右叶子
func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}