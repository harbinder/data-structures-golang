package tree

import "fmt"

/*
  Depth First Search
  - PreOrder:  root - left - right
  - InOrder:   left - root - right (gives sorted list for Binary Search Tree)
  - PostOrder: left - right - root

  Can be implemented using Stack Data Structure for Graphs
  Tree is also a type of graph called Binary Spanning Tree, which can be created
  out of Graph traversal using DFS
*/

func (bst *BinarySearchTree) PreOrderTraverse(root *BinaryTreeNode) {
	if root == nil {
		//fmt.Println("Leaf Node")
		return
	}

	fmt.Printf("%v->", root.Data)
	bst.PreOrderTraverse(root.Left)
	bst.PreOrderTraverse(root.Right)
}

func (bst *BinarySearchTree) InOrderTraverse(root *BinaryTreeNode) {
	if root == nil {
		//fmt.Println("Leaf Node")
		return
	}

	bst.InOrderTraverse(root.Left)
	fmt.Printf("%v->", root.Data)
	bst.InOrderTraverse(root.Right)
}

func (bst *BinarySearchTree) PostOrderTraverse(root *BinaryTreeNode) {
	if root == nil {
		//fmt.Println("Leaf Node")
		return
	}

	bst.PostOrderTraverse(root.Left)
	bst.PostOrderTraverse(root.Right)
	fmt.Printf("%v->", root.Data)
}
