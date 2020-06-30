package main

import (
	"algorithm/tree"
	"fmt"
	"math/rand"
	"time"
)

func test() {
	// var root1 *tree.AvlNode
	var root2 *tree.AvlNode
	arrInt := []int{5802, 7885, 9412, 4376, 667, 4362, 3069, 3535}
	for i := 0; i < len(arrInt); i++ {
		n := arrInt[i]
		// tree.TreeInsert(&root1, n, false)
		tree.TreeInsert(&root2, n, true)
		fmt.Println("++++++++++++++++++++++++++")
		tree.LevelTraverse(root2)
	}
	// tree.LevelTraverse(root1)
	// fmt.Println("++++++++++++++++++++++++++")
	// tree.LevelTraverse(root2)
}

func test2() {
	// var root1 *tree.AvlNode
	var root2 *tree.AvlNode
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 18; i++ {
		n := r.Intn(10000)
		// fmt.Printf("==%d", n)
		// tree.TreeInsert(&root1, n, false)
		tree.TreeInsert(&root2, n, true)
	}
	fmt.Printf("\n")
	// tree.LevelTraverse(root1)
	fmt.Println("++++++++++++++++++++++++++")
	tree.LevelTraverse(root2)
	var input int
	for {
		fmt.Scanf("%d", &input)
		if input == 0 {
			return
		}
		tree.TreeDelete(&root2, input)
		fmt.Println("++++++++++++++++++++++++++")
		tree.LevelTraverse(root2)

	}
}

func main() {
	test2()
}
