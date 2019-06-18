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
func (this *RBTree) leftRotate(n *node) {
	if n.right == nil {
		return
	}
	rn := n.right
	n.right = rn.left
	if rn.left != nil {
		n.left = rn.right
	}
	//TODO
}

//右旋
func (this *RBTree) rightRotate(n *node) {

}

//查找节点
func (this *RBTree) Get(value int) *node {
	return nil
}

//插入节点
func (this *RBTree) Insert(value int) {

}

//修复插入节点
func (this *RBTree) insertFixup(n *node) {

}

//删除节点
func (this *RBTree) Delete(value int) {

}

//修复删除节点
func (this *RBTree) deleteFixup() {

}
