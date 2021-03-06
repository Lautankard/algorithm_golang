package tree

import "fmt"

type AvlNode struct {
	Data  int
	Hight int //标记节点的高度
	Left  *AvlNode
	Right *AvlNode
}

func TreeHight(root *AvlNode) int {
	if root == nil {
		return 0
	} else {
		return Max(TreeHight(root.Left), TreeHight(root.Right)) + 1
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//节点的平衡因子
func TreeBalanceFactor(root *AvlNode) int {
	if root == nil {
		return 0
	} else {
		var left, right int
		if root.Left != nil {
			left = root.Left.Hight
		}
		if root.Right != nil {
			right = root.Right.Hight
		}
		return left - right
	}
}

// func TreeRotateLeft(root *AvlNode) *AvlNode {
// if root == nil || root.Right == nil {
// return root
// }
// right := root.Right
// root.Right = right.Left
// right.Left = root
//
// root.Hight = TreeHight(root)
// right.Hight = TreeHight(right)
//
// return nil
// }
//
// func TreeRotateRight(root *AvlNode) *AvlNode {
// if root == nil || root.Left == nil {
// return root
// }
// left := root.Left
// root.Left = left.Right
// left.Right = root
//
// root.Hight = TreeHight(root)
// left.Hight = TreeHight(left)
// return left
// }

// func TreeRebalance(root *AvlNode) {
// factor := TreeBalanceFactor(root)
// if factor > 1 && TreeBalanceFactor(root.Left) > 0 {
// root = TreeRotateRight(root)
// } else if factor > 1 && TreeBalanceFactor(root.Left) <= 0 {
// root.Left = TreeRotateLeft(root.Left)
// root = TreeRotateRight(root)
// } else if factor < -1 && TreeBalanceFactor(root.Right) <= 0 {
// root = TreeRotateLeft(root)
// } else if factor < -1 && TreeBalanceFactor(root.Right) > 0 {
// root.Right = TreeRotateRight(root.Right)
// root = TreeRotateLeft(root)
// }
// }

//左旋
func TreeRotateLeft(rootPtr **AvlNode) {
	root := *rootPtr
	if root == nil || root.Right == nil {
		return
	}
	right := root.Right
	root.Right = right.Left
	right.Left = root
	*rootPtr = right
	//旋转后需要重新计算被旋转节点的hight值
	root.Hight = TreeHight(root)
	right.Hight = TreeHight(right)
}

//右旋
func TreeRotateRight(rootPtr **AvlNode) {
	root := *rootPtr
	if root == nil || root.Left == nil {
		return
	}
	left := root.Left
	root.Left = left.Right
	left.Right = root
	*rootPtr = left
	root.Hight = TreeHight(root)
	left.Hight = TreeHight(left)
}

func TreeRebalance(rootPtr **AvlNode) {
	root := *rootPtr
	factor := TreeBalanceFactor(root)
	if factor > 1 && TreeBalanceFactor(root.Left) > 0 { //LL 左左：左高右低，数据插入到了左子树的左子树，只需要右旋
		TreeRotateRight(rootPtr)
	} else if factor > 1 && TreeBalanceFactor(root.Left) <= 0 { //LR, 左高右低，数据插在了左子树的右子树上，需要先左转再右转
		TreeRotateLeft(&root.Left)
		TreeRotateRight(rootPtr)
	} else if factor < -1 && TreeBalanceFactor(root.Right) <= 0 { //RR, 左低右高，数据插在了右子树的右子树，只需要左旋
		TreeRotateLeft(rootPtr)
	} else if factor < -1 && TreeBalanceFactor(root.Right) > 0 { //RL,左低右高，数据插在了右子树的左子树，需要先右旋再左旋
		TreeRotateRight(&root.Right)
		TreeRotateLeft(rootPtr)
	}
}

func TreeInsert(rootPtr **AvlNode, value int, b bool) {
	root := *rootPtr
	if root == nil {
		// fmt.Println("插入", value)
		root = &AvlNode{
			Data: value,
		}
		*rootPtr = root
	} else if root.Data == value {
		return
	} else if root.Data > value {
		TreeInsert(&root.Left, value, b)
	} else {
		TreeInsert(&root.Right, value, b)
	}
	//插入节点重新计算所影响节点的height
	root.Hight = TreeHight(root)
	if b {
		TreeRebalance(rootPtr)
	}
}

func FindMax(root *AvlNode) *AvlNode {
	if root == nil {
		return nil
	}
	if root.Right == nil {
		return root
	}
	return FindMax(root.Right)
}

func FindMin(root *AvlNode) *AvlNode {
	if root == nil {
		return nil
	}
	if root.Left == nil {
		return root
	}
	return FindMin(root.Left)
}

func TreeDelete(rootPtr **AvlNode, val int) {
	root := *rootPtr
	if root == nil {
		return
	}
	if val < root.Data {
		TreeDelete(&root.Left, val)
	} else if val > root.Data {
		TreeDelete(&root.Right, val)
	} else {
		if root.Right != nil {
			node := FindMin(root.Right)
			root.Data = node.Data
			TreeDelete(&root.Right, root.Data)
		} else if root.Left != nil {
			node := FindMax(root.Left)
			root.Data = node.Data
			TreeDelete(&root.Left, root.Data)
		} else {
			*rootPtr = nil
		}
	}

	root.Hight = TreeHight(root)
	TreeRebalance(rootPtr)
}

func MiddleTraverse(root *AvlNode) {
	if root == nil {
		return
	}
	MiddleTraverse(root.Left)
	fmt.Printf("-%d", root.Data)
	MiddleTraverse(root.Right)
}

func LevelTraverse(root *AvlNode) {
	if root == nil {
		return
	}
	var fifo []*AvlNode
	var tmp []*AvlNode
	fifo = append(fifo, root)

	for len(fifo) != 0 {
		r := fifo[0]
		fifo = fifo[1:]
		fmt.Printf("---%d+%d", r.Data, r.Hight)
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
		// if r.Left != nil {
		// fifo = append(fifo, r.Left)
		// }
		// if r.Right != nil {
		// fifo = append(fifo, r.Right)
		// }

	}
	fmt.Println()
}
