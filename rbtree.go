package rbtree

/*红黑树的性质：
1、每个节点都有一个颜色属性,每个节点或是红的或是黑的
2、根节点必须是黑的
3、每个叶子节点(nil节点)为黑
4、如果一个节点为红的,那么它的两个孩子都是黑的
5、每个节点到它子孙叶子节点的路径上的黑色节点个数是相同的*/

//节点颜色
const (
	RED   = 0
	BLACK = 1
)

//节点
type node struct {
	left   *node
	right  *node
	parent *node
	color  uint
	value  int
}

//红黑树
type RBTree struct {
	root  *node
	count uint
}

func NewRBTree() *RBTree {
	node := &node{left: nil, right: nil, parent: nil, color: BLACK, value: nil}
	return &RBTree{
		root:  node,
		count: 0,
	}
}

//左旋
/*
        |                        |
        X                        Y
     /      \                 /      \
    a        Y     ======>   X        c
           /   \          /    \
          b     c        a      b
*/
func (this *RBTree) leftRotate(x *node) {
	if x.right == nil {
		return
	}
	y := x.right
	x.right = y.left //Y的左节点变为X的右节点
	if x.right != nil { //???
		y.parent = x
	}
	y.parent = x.parent //Y节点的父节点为X节点的父节点
	if x.parent == nil { //如果X节点的父节点为nil，则Y就是根节点
		this.root = y
	} else if x == x.parent.left { //如果X节点为父节点的左节点，则将Y变为该父节点的左节点
		x.parent.left = y
	} else {
		x.parent.right = y
	}
	y.left = x //旋转，X成为Y的左节点，X的父节点为Y
	x.parent = y
}

//右旋
func (this *RBTree) rightRotate(n *node) {
	//TODO
}

//查找节点
func (this *RBTree) Get(value int) *node {
	//TODO
	return nil
}

//插入节点
func (this *RBTree) Insert(value int) {
	//TODO
}

//修复插入节点
func (this *RBTree) insertFixup(n *node) {
	//TODO
}

//删除节点
func (this *RBTree) Delete(value int) {
	//TODO
}

//修复删除节点
func (this *RBTree) deleteFixup() {
	//TODO
}
