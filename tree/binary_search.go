package tree

import (
	"fmt"
)

/*
  Node: memory location having DATA with left & right pointers
  Edge: connection/pointer from one NODE to other
  Tree: having NODES & EDGES
  SubTree: part of tree, which itself is a tree
  Left SubTree:
  Right SubTree:
	- Root:
	- Parent:
	- Child:
	- Siblings: nodes at same level
	- Leaf: having no child
	- Ancestors: nodes in the path towards root node
	- Descendents:
	- Depth: count of Edges when moving from node towards the root
	- Height: MAX. count of EDGES when moving from ROOT to LEAF NODE
	- Level: starts from 0
	Number of nodes @ level: 2^l
	root: l=0, total nodes: 2^0 = 1
		  l=1, total nodes: 2^1 = 2
		  l=2, total nodes: 2^2 = 4
	leaf: l=h, total nodes: 2^h
	total nodes in perfect tree:
		N = 2^0 + 2^1 + 2^2 ...... 2^h
		N = 2^(h+1) - 1
		2^(h+1) = N - 1
		h+1 = log(N-1)
		h = log(N)

	- Binary Tree (BT): in which every NODE has 0/1/2 CHILDREN
	- Strict BT: BT in which every NODE has exactly 2 CHILDREN
	- Complete BT: Strict BT  in which any NODE doesn't  have right NODE without left NODE
	- Perfect BT: Complete BT in which each NODE has left and right NODE except LEAF NODES
	- Balanced BT:
	- Binart Search Tree: in which every NODE has DATA > any node in its Left SubTree and Data < any node in its Right SubTree

	Time Complexity:
	  Operation       Array(unsorted)			Array(sorted)	LinkedList		BinartTree
	1. Search				O(n)					O(log n)		O(n)		O(log n)
	2. Insertion	O(1)-static,O(n)-dynamic 		O(n)			O(n)		O(log n)+
	3. Deletion     		O(n)					O(n)			O(n)		O(log n)+
*/

type BinaryTreeNode struct {
	Data  int
	Left  *BinaryTreeNode
	Right *BinaryTreeNode
}

type BinarySearchTree struct {
	Root   *BinaryTreeNode
	Count  int
	Height int
}

func (bst *BinarySearchTree) Create(data int) (newNode *BinaryTreeNode) {
	newNode = &BinaryTreeNode{Data: data}
	bst.Count++
	return
}

func (bst *BinarySearchTree) Insert(root *BinaryTreeNode, data int) *BinaryTreeNode {
	if root == nil {
		fmt.Printf("Insert node:%v", data)
		return bst.Create(data)
	}

	if root.Data > data {
		fmt.Println("Left")
		root.Left = bst.Insert(root.Left, data)
	} else {
		fmt.Println("Right")
		root.Right = bst.Insert(root.Right, data)
	}

	return root
}

func (bst *BinarySearchTree) Traverse(root *BinaryTreeNode) {
	if root == nil {
		fmt.Println("Leaf Node")
		return
	}

	fmt.Printf("\n%v\n", root.Data)
	//fmt.Printf("/\\\n")

	bst.Traverse(root.Left)
	bst.Traverse(root.Right)
}

func (bst *BinarySearchTree) Search(root *BinaryTreeNode, data int) (found string) {
	if root == nil {
		return "not found"

	}
	if data == root.Data {
		return "found"
	}

	if data > root.Data {
		found = bst.Search(root.Right, data)
	} else {
		found = bst.Search(root.Left, data)
	}
	return
}
