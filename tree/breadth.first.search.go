package tree

import (
	"container/list"
	"fmt"
)

/*
  Breadth First Search
  LevelOrder Traversal is used to traverse in BFS
*/
//var Qtree *list.List

func (bst *BinarySearchTree) LevelOrderTraverse(root *BinaryTreeNode) {

	if root == nil {
		return
	}

	Qtree := list.New()
	Qtree.PushBack(root)

	for Qtree.Len() > 0 {
		currentNode := Qtree.Front().Value.(*BinaryTreeNode)
		fmt.Printf("%v->", currentNode.Data)

		if currentNode.Left != nil {
			Qtree.PushBack(currentNode.Left)
		}
		if currentNode.Right != nil {
			Qtree.PushBack(currentNode.Right)
		}
		Qtree.Remove(Qtree.Front())
	}

	/*
		e := queue.Front() // First element
		fmt.Printf("%v\n", (e.Value))
		queue.Remove(e) // Dequeue
	*/
	return
}
