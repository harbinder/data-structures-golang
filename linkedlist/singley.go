package linkedlist

import (
	"fmt"
	"strconv"
)

type Node struct {
	Data int
	Next *Node
}

func Create(data int) (node *Node) {
	node = &Node{
		Data: data,
	}
	return
}

func Traverse(head *Node) {
	current := head
	for {
		if current.Next != nil {
			fmt.Printf("%v->", current.Data)
		} else {
			fmt.Printf("%v\n", current.Data)
			return
		}
		current = current.Next
	}
}

func AddOld(data int, position string, head *Node) (*Node, error) {
	newNode := Create(data)
	switch position {
	case "start":
		newNode.Next = head
		head = newNode
	case "end":
		current := head
		for {
			if current.Next == nil {
				current.Next = newNode
				return head, nil
			}
			current = current.Next
		}
	default:
		current := head
		currPos := 1
		for {
			pos, _ := strconv.Atoi(position)
			if currPos == pos-1 {
				newNode.Next = current.Next
				current.Next = newNode
				return head, nil
			}
			if current.Next == nil {
				err := fmt.Errorf("Position not found")
				return head, err
			}
			current = current.Next
		}
	}
	return head, nil
}

func Add(data int, position int, head *Node) (*Node, error) {
	current := head
	prev := current
	currPos := 1
	var err error
	for {
		if position == currPos {
			newNode := Create(data)
			newNode.Next = current
			if position == 1 {
				head = newNode
			} else {
				prev.Next = newNode
			}
			return head, nil
		}
		if current.Next == nil {
			err = fmt.Errorf("Position not found")
			return head, err
		}
		prev = current
		current = current.Next
		currPos++
	}

	return head, nil
}

func Delete(data int, head *Node) (*Node, error) {
	var err error
	current := head
	prev := head
	posCount := 1
	for {
		if current.Data == data {
			if posCount == 1 {
				head = current.Next
			} else {
				prev.Next = current.Next
			}

			// TODO: delete node
			current = nil
			return head, nil
		}
		if current.Next == nil {
			err = fmt.Errorf("Data not found")
			return head, err
		}

		prev = current
		current = current.Next
		posCount++
	}
}

func Reverse(head *Node) *Node {
	fmt.Println("\nReserve List")

	current := head
	next := current.Next
	var prev *Node

	for {

		current.Next = prev
		prev = current
		current = next

		if next != nil {
			next = next.Next
		}

		if current == nil {
			return prev
		}

	}
}
