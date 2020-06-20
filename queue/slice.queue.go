package queue

import (
	"fmt"

	"my.datastructure/tree"
)

type (
	TreeQ  []*tree.BinaryTreeNode
	SliceQ []*int
)

func (q TreeQ) Enqueue(nodePtr *tree.BinaryTreeNode) {
	if q == nil {
		q = make(TreeQ, 10)
	}

	q = append(q, nodePtr)
}

func (q TreeQ) Dequeue() (nodePtr *tree.BinaryTreeNode) {
	if len(q) != 0 {
		nodePtr = q[0]
		q = q[1:]
	}
	return
}

func Enqueue(q []*int, data int) {
	// if q == nil {
	// 	q = make(SliceQ, 10)
	// }
	q = append(q, &data)
	fmt.Printf("enque %v:%v\n", len(q), *q[0])
}

func Dequeue(q []*int) (data int) {
	if len(q) != 0 {
		data = *q[0]
		q = q[1:]
		fmt.Printf("deque length: %v\n", len(q))
	}
	return
}

func Traverse(q []*int) {
	for k, v := range q {
		fmt.Printf("%v -> %p ", k, v)
	}
}

func (q SliceQ) Enqueue(data int) {
	// if q == nil {
	// 	q = make(SliceQ, 10)
	// }
	q = append(q, &data)
	fmt.Printf("enque %v:%v\n", len(q), *q[0])
}

func (q SliceQ) Dequeue() (data int) {
	if len(q) != 0 {
		data = *q[0]
		q = q[1:]
		fmt.Printf("deque length: %v\n", len(q))
	}
	return
}

func (q SliceQ) Traverse() {
	for k, v := range q {
		fmt.Printf("%v -> %p ", k, v)
	}
}
