package tree

import "fmt"

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
