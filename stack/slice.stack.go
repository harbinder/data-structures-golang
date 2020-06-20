package stack

import (
	"fmt"
)

type StackArr struct {
	top   int
	count int
	arr   []int
	//next  *StackArr
}

func NewStack() *StackArr {
	return &StackArr{
		top: -1,
	}
}
func (sn *StackArr) isEmpty() bool {
	if sn.top == -1 {
		return true
	}
	return false
}

func (sn *StackArr) isFull() bool {
	if sn.top == sn.count-1 {
		return true
	}
	return false
}

func (sn *StackArr) Push(data int) {
	// if sn.isFull() {
	// 	fmt.Println("stack overflow")
	// 	return
	// }

	// if len(sn.arr) == 0 {
	// 	sn.arr = make([]int, 0)
	// }
	sn.arr = append(sn.arr, data)

	sn.top++
	sn.count++
	fmt.Printf("\nPUSH - data: %v, top: %v , count: %v\n", data, sn.top, sn.count)
}

func (sn *StackArr) Pop() (data int) {
	if sn.isEmpty() {
		fmt.Println("stack underflow")
		return
	}

	data = sn.arr[sn.top]
	sn.arr[sn.top] = 0

	sn.top--
	sn.count--

	fmt.Printf("\nPOP - data: %v, top: %v , count: %v\n", data, sn.top, sn.count)

	return
}
