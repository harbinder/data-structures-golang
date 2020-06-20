package stack

import (
	"fmt"
)

type StackNode struct {
	data int
	next *StackNode
}

type StackList struct {
	head *StackNode
}

func NewStackList() *StackList {
	return &StackList{}
}

func (sn *StackList) createNode(data int) *StackNode {
	return &StackNode{
		data: data,
	}
}
func (sn *StackList) isEmpty() (e bool) {
	if sn.head == nil {
		e = true
	}
	return
}

func (sn *StackList) Push(data int) {
	newNode := sn.createNode(data)

	if sn.head != nil {
		newNode.next = sn.head
	}
	sn.head = newNode
	fmt.Println("PUSH - data: ", data)
}

func (sn *StackList) Pop() (data int) {

	if sn.isEmpty() {
		fmt.Println("stack underflow")
		return
	}

	data = sn.head.data
	//deleteNode := sn.head
	sn.head = sn.head.next

	fmt.Println("POP - data: ", data)
	return
}
