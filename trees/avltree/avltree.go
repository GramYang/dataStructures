package avltree

import (
	"dataStructures/utils"
	"fmt"
)

type Tree struct {
	Root       *Node
	Comparator utils.Comparator
	size       int
}

type Node struct {
	Key      interface{}
	Value    interface{}
	Parent   *Node    // Parent node
	Children [2]*Node // Children nodes
	b        int8
}

//用comparator实现一个avl树实例
func NewWith(comparator utils.Comparator) *Tree {
	return &Tree{Comparator: comparator}
}

//返回一个avl树实例，avl树的节点都是int值
func NewWithIntComparator() *Tree {
	return &Tree{Comparator: utils.IntComparator}
}

//返回avl树的节点都是string的实例
func NewWithStringComparator() *Tree {
	return &Tree{Comparator: utils.StringComparator}
}

//存入一个节点
func (t *Tree) Put(key interface{}, value interface{}) {
	t.put(key, value, nil, &t.Root)
}

//获取一个节点
func (t *Tree) Get(key interface{}) (value interface{}, found bool) {
	n := t.Root
	for n != nil {
		cmp := t.Comparator(key, n.Key)
		switch {
		case cmp == 0:
			return n.Value, true
		case cmp < 0:
			n = n.Children[0]
		case cmp > 0:
			n = n.Children[1]
		}
	}
	return nil, false
}

//删除某个节点
func (t *Tree) Remove(key interface{}) {
	t.remove(key, &t.Root)
}

//判断avl树是否为空
func (t *Tree) Empty() bool {
	return t.size == 0
}

//返回avl树的节点数
func (t *Tree) Size() int {
	return t.size
}

//返回avl树中所有节点的key值的切片，切片值升序
func (t *Tree) Keys() []interface{} {
	keys := make([]interface{}, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		keys[i] = it.Key()
	}
	return keys
}

//返回avl树中所有节点的value值的切片，切片之升序
func (t *Tree) Values() []interface{} {
	values := make([]interface{}, t.size)
	it := t.Iterator()
	for i := 0; it.Next(); i++ {
		values[i] = it.Value()
	}
	return values
}

//返回avl树的最左叶子，即最小值。树为空则返回nil
func (t *Tree) Left() *Node {
	return t.bottom(0)
}

//返回avl树的最右叶子，即最大值。树为空则返回nil
func (t *Tree) Right() *Node {
	return t.bottom(1)
}

//返回Floor节点，Floor节点是仅小于或者等于key节点的节点
func (t *Tree) Floor(key interface{}) (floor *Node, found bool) {
	found = false
	n := t.Root
	for n != nil {
		c := t.Comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			n = n.Children[0]
		case c > 0:
			floor, found = n, true
			n = n.Children[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

//返回Ceiling节点，仅大于或等于key节点的节点
func (t *Tree) Ceiling(key interface{}) (floor *Node, found bool) {
	found = false
	n := t.Root
	for n != nil {
		c := t.Comparator(key, n.Key)
		switch {
		case c == 0:
			return n, true
		case c < 0:
			floor, found = n, true
			n = n.Children[0]
		case c > 0:
			n = n.Children[1]
		}
	}
	if found {
		return
	}
	return nil, false
}

//清空avl树
func (t *Tree) Clear() {
	t.Root = nil
	t.size = 0
}

//avl树输出成字符串
func (t *Tree) String() string {
	str := "AVLTree\n"
	if !t.Empty() {
		output(t.Root, "", true, &str)
	}
	return str
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Key)
}

//向avl树中存入节点，存完后还会调整成avl树
func (t *Tree) put(key interface{}, value interface{}, p *Node, qp **Node) bool {
	q := *qp
	if q == nil { //说明树是空的，直接将存入节点作为父节点，put成功
		t.size++
		*qp = &Node{Key: key, Value: value, Parent: p}
		return true
	}
	c := t.Comparator(key, q.Key)
	if c == 0 { //和节点相等，put失败
		q.Key = key
		q.Value = value
		return false
	}
	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	var fix bool
	fix = t.put(key, value, q, &q.Children[a]) //比节点小则递归左叶子，比节点大则递归右叶子
	if fix {
		return putFix(int8(c), qp) //avl树插入调整
	}
	return false
}

func (t *Tree) remove(key interface{}, qp **Node) bool {
	q := *qp
	if q == nil {
		return false
	}
	c := t.Comparator(key, q.Key)
	if c == 0 {
		t.size--
		if q.Children[1] == nil {
			if q.Children[0] != nil {
				q.Children[0].Parent = q.Parent
			}
			*qp = q.Children[0]
			return true
		}
		fix := removeMin(&q.Children[1], &q.Key, &q.Value)
		if fix {
			return removeFix(-1, qp)
		}
		return false
	}
	if c < 0 {
		c = -1
	} else {
		c = 1
	}
	a := (c + 1) / 2
	fix := t.remove(key, &q.Children[a])
	if fix {
		return removeFix(int8(-c), qp)
	}
	return false
}

func removeMin(qp **Node, minKey *interface{}, minVal *interface{}) bool {
	q := *qp
	if q.Children[0] == nil {
		*minKey = q.Key
		*minVal = q.Value
		if q.Children[1] != nil {
			q.Children[1].Parent = q.Parent
		}
		*qp = q.Children[1]
		return true
	}
	fix := removeMin(&q.Children[0], minKey, minVal)
	if fix {
		return removeFix(1, qp)
	}
	return false
}

//avl树插入调整，t为根节点指针，c=-1表示插入左子树，c=1表示插入右子树
func putFix(c int8, t **Node) bool {
	s := *t
	if s.b == 0 {//上一次刚进行了旋转
		s.b = c
		return true
	}
	if s.b == -c {
		s.b = 0
		return false
	}
	if s.Children[(c+1)/2].b == c {
		s = singlerot(c, s) //单向旋转
	} else {
		s = doublerot(c, s) //双向旋转
	}
	*t = s
	return false
}

func removeFix(c int8, t **Node) bool {
	s := *t
	if s.b == 0 {
		s.b = c
		return false
	}
	if s.b == -c {
		s.b = 0
		return true
	}
	a := (c + 1) / 2
	if s.Children[a].b == 0 {
		s = rotate(c, s)
		s.b = -c
		*t = s
		return false
	}
	if s.Children[a].b == c {
		s = singlerot(c, s)
	} else {
		s = doublerot(c, s)
	}
	*t = s
	return true
}

//单旋转
func singlerot(c int8, s *Node) *Node {
	s.b = 0
	s = rotate(c, s)
	s.b = 0
	return s
}

//双旋转
func doublerot(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = rotate(-c, s.Children[a])
	p := rotate(c, s)

	switch {
	default:
		s.b = 0
		r.b = 0
	case p.b == c:
		s.b = -c
		r.b = 0
	case p.b == -c:
		s.b = 0
		r.b = c
	}

	p.b = 0
	return p
}

func rotate(c int8, s *Node) *Node {
	a := (c + 1) / 2
	r := s.Children[a]
	s.Children[a] = r.Children[a^1]
	if s.Children[a] != nil {
		s.Children[a].Parent = s
	}
	r.Children[a^1] = s
	r.Parent = s.Parent
	s.Parent = r
	return r
}

//d=1遍历到最右叶子节点后返回。d=0遍历最左叶子节点后返回
func (t *Tree) bottom(d int) *Node {
	n := t.Root
	if n == nil {
		return nil
	}
	for c := n.Children[d]; c != nil; c = n.Children[d] {
		n = c
	}
	return n
}

//遍历到离n节点最近的小于n的节点
func (n *Node) Prev() *Node {
	return n.walk1(0)
}

//遍历到离节点n最近的大于n的节点
func (n *Node) Next() *Node {
	return n.walk1(1)
}

//a=1，遍历到离节点n最近的大于n的节点。a=0，遍历到离节点n最近的小于n的节点。这个节点可能是子节点，也可能是父节点。
func (n *Node) walk1(a int) *Node {
	if n == nil {
		return nil
	}
	if n.Children[a] != nil {
		n = n.Children[a]
		for n.Children[a^1] != nil {
			n = n.Children[a^1]
		}
		return n
	}
	p := n.Parent
	for p != nil && p.Children[a] == n {
		n = p
		p = p.Parent
	}
	return p
}

//以树型图的方式输出avl树
func output(node *Node, prefix string, isTail bool, str *string) {
	if node.Children[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		output(node.Children[1], newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Children[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		output(node.Children[0], newPrefix, true, str)
	}
}