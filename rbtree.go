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
    a        Y    =======>   X        c
           /   \          /    \
          b     c        a      b
*/
func (this *RBTree) leftRotate(x *node) {
	if x.right == nil {
		return
	}
	y := x.right        //Y为X的左孩子
	x.right = y.left    //X的右节点为Y的左节点
	y.left.parent = x   //Y节点的左节点的父节点为X
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
/*
            |                    |
            Y                    X
         /      \             /      \
        X        c  ======>  a        Y
     /    \                         /   \
    a      b                       b     c
*/
func (this *RBTree) rightRotate(y *node) {
	if y.left == nil {
		return
	}
	x := y.left         //X为Y的左孩子
	y.left = x.right    //Y的左节点为X的右节点
	x.right.parent = y  //X右节点的父节点为Y
	x.parent = y.parent //X的父节点为Y的父节点
	if y.parent == nil { //如果Y的父节点为空，那么X就为根节点
		this.root = x
	} else if y == y.parent.left { //如果Y为Y父节点的左节点，那么Y父节点的左节点为X
		y.parent.left = x
	} else {
		y.parent.right = x
	}
	x.right = y //旋转，X的右节点为Y，Y的父节点为X
	y.parent = x
}

//查找节点
func (this *RBTree) Get(value int) *node {
	n := this.root
	for n != nil {
		if value < n.value { //小于当前节点就往左找
			n = n.left
		} else if value > n.value { //大于当前节点往右找
			n = n.right
		} else {
			return n
		}
	}
	return nil
}

//插入节点
func (this *RBTree) Insert(v int) {
	if this.root == nil { //没有根节点则当前节点为根节点
		this.root = &node{color: BLACK, value: v,}
	}
	n := this.root
	insertNode := &node{color: RED, value: v}
	var insertNodeParent *node //插入节点的父节点
	for n != nil { //确定插入节点的父节点
		insertNodeParent = n
		if v < n.value {
			n = n.left
		} else if v > n.value {
			n = n.right
		} else {
			//插入相同的值
			return
		}
	}
	insertNode.parent = insertNodeParent //设置新节点的父节点
	if v < insertNodeParent.value {
		insertNodeParent.left = insertNode
	} else {
		insertNodeParent.right = insertNode
	}
	this.insertFixup(insertNode)
}

//修复插入节点(核心思路:将红色的节点移到根节点，然后将根节点设为黑色)
/*
1、当前节点的父节点是红色，且当前节点的叔叔节点也为红色
   处理方式：将父节点和叔叔节点设置为黑色，将祖父节点设置为红色并为当前节点，继续对当前节点进行操作
2、当前节点的父节点为红色，叔叔节点为黑色，且当前节点为父节点的右节点
   处理方式：将父节点作为新的当前节点，对新的当前节点为支点进行左旋
3、当前节点的父节点为红色，叔叔节点为黑色，且当前节点为父节点的左节点
   处理方式：将父节点设为黑色，祖父节点设为红色，以祖父节点为支点进行右旋
*/
func (this *RBTree) insertFixup(n *node) {
	for n.parent.color == RED {
		if n.parent == n.parent.parent.left { //当前节点的父节点为祖父节点的左孩子
			uncleNode := n.parent.parent.right
			if uncleNode.color == RED { //case 1
				n.parent.color = BLACK
				uncleNode.color = BLACK
				n.parent.parent.color = RED
				n = n.parent.parent
			} else { //叔叔节点为黑色
				if uncleNode == uncleNode.parent.left { //case 3
					uncleNode.parent.color = BLACK
					uncleNode.parent.parent.color = RED
					this.rightRotate(uncleNode.parent.parent)
				} else { //case 2
					uncleNode = uncleNode.parent
					this.leftRotate(uncleNode)
				}
			}
		} else {
			uncleNode := n.parent.parent.left
			if uncleNode.color == RED { //case 1
				n.parent.color = BLACK
				uncleNode.color = BLACK
				n.parent.parent.color = RED
				n = n.parent.parent
			} else { //叔叔节点为黑色
				if uncleNode == uncleNode.parent.left { //case 3
					uncleNode.parent.color = BLACK
					uncleNode.parent.parent.color = RED
					this.rightRotate(uncleNode.parent.parent)
				} else { //case 2
					uncleNode = uncleNode.parent
					this.leftRotate(uncleNode)
				}
			}
		}
	}
	this.root.color = BLACK
}

//删除节点
func (this *RBTree) Delete(value int) {
	//TODO
}

//修复删除节点
func (this *RBTree) deleteFixup() {
	//TODO
}
