package tree

import "fmt"

/*
原则：
1.节点不为红色即为黑色
2.根节点为黑色
3.叶子节点（nil节点）为黑色
4.如果节点为红色，则其子节点为黑色
5.任何一个节点起到叶子节点的每条路径含有相同数目的黑色节点
*/

type ColorType int

const (
	Black ColorType = iota + 1
	Red
)

func (color ColorType) IsRed() bool {
	return color == Red
}

func (color ColorType) IsBlack() bool {
	return color == Black
}

type RBTreeInsertFunc func(root, node, sentinel *RBNode)

type RBNode struct {
	Data   int
	Color  ColorType
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
}

type RBTree struct {
	Root     *RBNode
	Sentinel *RBNode
	Insert   RBTreeInsertFunc
}

func (rbt *RBTree) InsertValue(value int) {
	node := &RBNode{
		Data:  value,
		Color: Red,
	}
	rbt.InsertNode(node)
}

func (rbt *RBTree) InsertNode(node *RBNode) {
	rootPtr := &rbt.Root
	root := rbt.Root
	if root == nil {
		rbt.Root = node
		node.Parent = nil
		node.Left = rbt.Sentinel
		node.Right = rbt.Sentinel
		node.Color = Black
		return
	}
	rbt.Insert(root, node, rbt.Sentinel)

	for node != root && node.Parent.Color.IsRed() {
		if node.Parent == node.Parent.Parent.Left {
			uncle := node.Parent.Parent.Right
			if uncle.Color.IsRed() {
				node.Parent.Color = Black
				uncle.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Right {
					node = node.Parent
					rbt.LeftRotate(node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				rbt.RightRotate(node.Parent.Parent)
			}
		} else {
			uncle := node.Parent.Parent.Left
			if uncle.Color.IsRed() {
				uncle.Color = Black
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				node = node.Parent.Parent
			} else {
				if node == node.Parent.Left {
					node = node.Parent
					rbt.RightRotate(node)
				}
				node.Parent.Color = Black
				node.Parent.Parent.Color = Red
				rbt.LeftRotate(node.Parent.Parent)
			}
		}
	}
	(*rootPtr).Color = Black

}

func (rbt *RBTree) Delete(val int) {
	node := rbt.Root
	if node == nil || node == rbt.Sentinel {
		return
	}
	for node != rbt.Sentinel {
		if val < node.Data {
			node = node.Left
		} else if val > node.Data {
			node = node.Right
		} else {
			break
		}
	}
	if node != rbt.Sentinel {
		rbt.DeleteNode(node)
	}
}

func (rbt *RBTree) DeleteNode(node *RBNode) {
	var temp, subst *RBNode
	if node.Left == rbt.Sentinel {
		temp = node.Left
		subst = node
	} else if node.Right == rbt.Sentinel {
		temp = node.Left
		subst = node
	} else {
		subst = rbt.FindMin(node.Right)
		temp = subst.Right
	}
	if subst == rbt.Root {
		rbt.Root = temp
		temp.Color = Black
		temp.Parent = nil
		node.Left = nil
		node.Right = nil
		node.Parent = nil
		return
	}

	isRed := (subst.Color == Red)
	if subst == subst.Parent.Left {
		subst.Parent.Left = temp
	} else {
		subst.Parent.Right = temp
	}
	// if node != subst.Parent {
	if temp != rbt.Sentinel {
		temp.Parent = subst.Parent
	}
	// }
	if node != subst {
		subst.Left = node.Left
		subst.Right = node.Right
		subst.Parent = node.Parent
		if node == rbt.Root {
			rbt.Root = subst
		} else if node == node.Parent.Left {
			node.Parent.Left = subst
		} else {
			node.Parent.Right = subst
		}

		if subst.Left != rbt.Sentinel {
			subst.Left.Parent = subst
		}
		if subst.Right != rbt.Sentinel {
			subst.Right.Parent = subst
		}
	}
	node.Left = nil
	node.Right = nil
	node.Parent = nil

	if isRed {
		return
	}

	for temp != rbt.Root && temp.Color == Black { //如果temp为红色，跳出循环直接temp改为黑色
		if temp == temp.Parent.Left {
			bro := temp.Parent.Right
			if bro.Color == Red { //不改变此子树的情况下，构造bro为黑色
				bro.Color = Black
				temp.Parent.Color = Red
				rbt.LeftRotate(temp.Parent)
				bro = temp.Parent.Right
			}
			//bro为黑色
			if bro.Left.Color == Black && bro.Right.Color == Black { //子节点都为黑色，bro设置为红色，两边黑色数目相等
				bro.Color = Red
				temp = temp.Parent
			} else {
				//构造近临temp的bro的子树为黑色，远离的为红色，用于后期将远离的红色变为黑色
				if bro.Right.Color == Black { //右黑左红需要转化为左黑右红
					bro.Left.Color = Black
					bro.Color = Red
					rbt.RightRotate(bro)
					bro = temp.Parent.Right
				}

				/*左黑右红 此时分为两种情况，但是处理都一样
				第一种  parent(黑）                    bro(黑)
						 /    \                         /  \
					    /      \                       /    \
					temp(黑)  bro(黑） ------->   paretn(黑）br(黑)
					           /  \                  /  \       \
							  /    \                /    \       \
							bl(黑） br(红)      temp(黑) bl(黑)  (黑)
							          \
									   \
									  (黑)

				第二种  parent(红)                    bro(红)
						 /    \                         /  \
					    /      \                       /    \
					temp(黑)  bro(黑） ------->   paretn(黑）br(黑)
					           /  \                  /  \      \
							  /    \                /    \      \
							bl(黑） br(红)      temp(黑) bl(黑) (黑)
							          \
									   \
									   (黑)
				*/
				bro.Color = temp.Parent.Color
				temp.Color = Black
				bro.Right.Color = Black
				rbt.LeftRotate(temp.Parent)
				temp = rbt.Root
			}
		} else {
			bro := temp.Parent.Left
			if bro.Color == Red {
				bro.Color = Black
				temp.Parent.Color = Red
				rbt.RightRotate(temp.Parent)
				bro = temp.Parent.Left
			}
			if bro.Left.Color == Black && bro.Right.Color == Black {
				bro.Color = Red
				temp = temp.Parent
			} else {
				if bro.Left.Color == Black {
					bro.Right.Color = Black
					bro.Color = Red
					rbt.LeftRotate(bro)
					bro = temp.Parent.Left
				}
				bro.Color = temp.Parent.Color
				bro.Left.Color = Black
				temp.Parent.Color = Black
				rbt.RightRotate(temp.Parent)
				temp = rbt.Root
			}
		}
	}
	temp.Color = Black
}

func (rbt *RBTree) FindMin(node *RBNode) *RBNode {
	for node.Left != rbt.Sentinel {
		node = node.Left
	}
	return node
}

func (rbt *RBTree) LevelTraverse() {
	root := rbt.Root
	if root == nil {
		return
	}
	var fifo []*RBNode
	var tmp []*RBNode
	fifo = append(fifo, root)

	for len(fifo) != 0 {
		r := fifo[0]
		fifo = fifo[1:]
		fmt.Printf("---%d+%d", r.Data, r.Color)
		if r.Left != nil {
			tmp = append(tmp, r.Left)
		}
		if r.Right != nil {
			tmp = append(tmp, r.Right)
		}
		if len(fifo) == 0 && len(tmp) > 0 {
			fmt.Printf("\n")
			fifo = tmp[0:len(tmp)]
			tmp = tmp[len(tmp):]
			fmt.Println("----------------------")
		}
	}
	fmt.Println()
}

func (rbt *RBTree) LeftRotate(node *RBNode) {
	temp := node.Right
	node.Right = temp.Left
	if temp.Left != rbt.Sentinel {
		temp.Left.Parent = node
	}
	temp.Left = node
	temp.Parent = node.Parent
	if node == rbt.Root {
		rbt.Root = temp
	} else if node == node.Parent.Left {
		node.Parent.Left = temp
	} else if node == node.Parent.Right {
		node.Parent.Right = temp
	}
	node.Parent = temp

}

func (rbt RBTree) RightRotate(node *RBNode) {
	temp := node.Left
	node.Left = temp.Right
	if temp.Right != rbt.Sentinel {
		temp.Right.Parent = node
	}
	temp.Right = node
	temp.Parent = node.Parent
	if node == rbt.Root {
		rbt.Root = temp
	} else if node == node.Parent.Left {
		node.Parent.Left = temp
	} else {
		node.Parent.Right = temp
	}
	node.Parent = temp
}

func DefaultInsert(parent *RBNode, node, sentinel *RBNode) {
	var p **RBNode
	for {
		if parent.Data > node.Data {
			p = &parent.Left
		} else if parent.Data < node.Data {
			p = &parent.Right
		} else {
			return
		}
		if *p == sentinel {
			break
		}
		parent = *p

	}
	*p = node
	node.Parent = parent
	node.Left = sentinel
	node.Right = sentinel
	node.Color = Red

}

func NewDefaultRBTree() *RBTree {
	sentinel := &RBNode{
		Color: Black,
	}
	return &RBTree{
		Sentinel: sentinel,
		Insert:   DefaultInsert,
	}
}

// func RBTreeRotateLeft(rootPtr **RBNode, node, sentinel *RBNode) {
// temp := node.Right
// node.Right = temp.Left
// if temp.Left != sentinel {
// temp.Left.Parent = node
// }
// temp.Left = node
// temp.Parent = node.Parent
// if node == *rootPtr {
// *rootPtr = temp
// } else if node == node.Parent.Left {
// node.Parent.Left = temp
// } else if node == node.Parent.Right {
// node.Parent.Right = temp
// }
// node.Parent = temp
// }
//
// func RBTreeRotateRight(rootPtr **RBNode, node, sentinel *RBNode) {
// temp := node.Left
// node.Left = temp.Right
// if temp.Right != sentinel {
// temp.Right.Parent = node
// }
// temp.Right = node
// temp.Parent = node.Parent
// if node == *rootPtr {
// *rootPtr = temp
// } else if node == node.Parent.Left {
// node.Parent.Left = temp
// } else {
// node.Parent.Right = temp
// }
// node.Parent = temp
// }
