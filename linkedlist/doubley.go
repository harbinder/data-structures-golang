package linkedlist

import (
	"fmt"
)

type DoubleyNode struct {
	Data int
	Next *DoubleyNode
	Prev *DoubleyNode
}

type DoubleyList struct {
	Head   *DoubleyNode
	Tail   *DoubleyNode
	Length int
}

func (dl *DoubleyList) Create(data int) (dn *DoubleyNode) {

	dn = &DoubleyNode{
		Data: data,
	}
	return
}

func (dl *DoubleyList) IndexOf(data int) (index int, err error) {
	traversal := "forward"
	current := dl.Head
	var currentPos int
	for {

		if data == current.Data {
			if traversal == "forward" {
				index = currentPos
			} else {
				index = (dl.Length - 1) - currentPos
			}
			return
		}

		if traversal == "forward" {
			if current.Next == nil {
				err = fmt.Errorf("Value not available")
				return
			}
			current = current.Next
		} else {
			if current.Prev == nil {
				err = fmt.Errorf("Value not available")
				return
			}
			current = current.Prev
		}
		currentPos++
	}
}

func (dl *DoubleyList) AtIndex(index int) (data int, err error) {
	if index > dl.Length-1 {
		err = fmt.Errorf("Index out of bound")
		return
	}

	// check whether start to traversal from head or tail
	var currentPos int
	current := dl.Head

	traversal := "forward"
	if index > dl.Length/2 {
		current = dl.Tail
		currentPos = dl.Length - 1
		traversal = "backward"
	}
	fmt.Println("Traversing ", traversal)

	for {

		if index == currentPos {
			data = current.Data
			return
		}

		if traversal == "forward" {
			if current.Next == nil {
				err = fmt.Errorf("forward: Index out of bound")
				return
			}
			current = current.Next
		} else {
			if current.Prev == nil {
				err = fmt.Errorf("backward: Index out of bound")
				return
			}
			current = current.Prev
		}

		currentPos++
	}
}

func (dl *DoubleyList) Traversal() {
	fmt.Printf("\nLength: %v\n", dl.Length)

	current := dl.Head
	for {
		if dl.Length == 0 {
			return
		}

		fmt.Printf("%v->", current.Data)
		if current.Next == nil {
			fmt.Println()
			return
		}

		current = current.Next
	}
}

func (dl *DoubleyList) Add(data, index int) {
	current := dl.Head
	var prev *DoubleyNode
	var posCounter int

	for {

		if index == posCounter {
			newNode := dl.Create(data)

			if dl.Length == 0 {
				fmt.Println("adding the first node")
				dl.Head = newNode
				dl.Tail = newNode
			} else if index == 0 {
				fmt.Println("adding at the start of list")
				dl.Head = newNode
				newNode.Next = current
				current.Prev = newNode
			} else {
				fmt.Println("adding at any other index except last")
				newNode.Prev = current.Prev
				newNode.Next = prev.Next
				prev.Next = newNode
				current.Prev = newNode
			}

			dl.Length++
			return
		}

		if current.Next == nil {
			fmt.Println("adding at the end of list")

			newNode := dl.Create(data)

			dl.Tail = newNode
			current.Next = newNode
			newNode.Prev = current

			dl.Length++
			return
		}

		prev = current
		current = current.Next
		posCounter++
	}

}
