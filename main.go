package main

import (
	"fmt"
	"runtime"
	"time"

	"my.datastructure/linkedlist"
	"my.datastructure/queue"
	"my.datastructure/sort"
	"my.datastructure/stack"
	"my.datastructure/tree"
)

func main() {
	fmt.Println("Data Structures ........")

	//Sort()
	//StackList()
	//StackArr()

	//Queue()

	//SinglyList()
	//DoublyList()

	BinarySearchTree()
}

func Sort() {
	fmt.Println("Sort List....")
	//QuickSort()
	//MergeSort() // recursive merge sort
	//TwoWayMerge() // iterative merge sort
	ListMerge() // Merging 2 LinkedList
}

func ListMerge() {
	// create 2 Linked Lists
	list1 := sort.NewMergeList()
	slice1 := []int{2, 4, 6, 8, 10}
	for _, v := range slice1 {
		list1.InsertNode(v)
	}

	list2 := sort.NewMergeList()
	slice2 := []int{1, 3, 5, 7, 9}
	for _, v := range slice2 {
		list2.InsertNode(v)
	}

	list1.TraverseList()
	list2.TraverseList()

	// merge both linked lists
	newHead := sort.MergeLinkedLists(list1.Head, list2.Head)
	mergedList := &sort.MergeList{Head: newHead}
	mergedList.TraverseList()

}

func TwoWayMerge() {
	fmt.Println("Two-Way Merge Sort ....")
	unsortedList := []int{7, 3, 5, 1, 6, 0, 2, 4}
	//[7,3,9,5,1,6,0,2,8,4]
	//[3 7 5 9 1 6 0 2 4 8]
	ms := &sort.MergeSortStruct{
		List: unsortedList,
	}
	sortedList := ms.TwoWayMerge()
	fmt.Println(sortedList)
}

func MergeSort() {
	fmt.Println("Merge Sort ....")

	/*
	   Analysis: Recursion vs Goroutine

	   FOR Lesser Elements
	   - Resursion taking less time
	   - Goroutines taking more time

	   100M ELEMENTS
	   R: takes 25.64s
	   G: takes 23.89s
	*/

	flag := "concurrent" // "concurrent" | "recursive"
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("runtime.NumCPU(): ", runtime.NumCPU())

	var unsortedList, sortedList []int
	//unsortedList := []int{7, 3, 9, 5, 1, 6, 0, 2, 8, 4}
	for i := 0; i < 100; i++ {
		unsortedList = append(unsortedList, i)
	}
	ms := &sort.MergeSortStruct{
		List: unsortedList,
	}

	tStart := time.Now()
	switch flag {
	case "recursive":
		sortedList = ms.MergeSort(0, len(unsortedList)-1)
	case "concurrent":
		sortedList = ms.MergeSort(0, len(unsortedList)-1)
	}
	tEnd := time.Now()
	duration := tEnd.Sub(tStart)

	fmt.Println("Duration: ", duration)
	fmt.Println(len(sortedList))
}

func Merge() {

	fmt.Println("Merging of Lists....")
	firstList := []int{1, 3, 5, 7, 9}
	secondList := []int{2, 4, 6, 8, 10}
	m := sort.MergeStruct{
		FirstSortedList:  firstList,
		SecondSortedList: secondList,
	}
	mergedList := m.Merge()
	fmt.Println(mergedList)
}

func QuickSort() {

	fmt.Println("Quick Sort List....")
	arr := []int{10, 16, 8, 12, 15, 6, 3, 9, 5, 999}
	qs := sort.QuickSortStruct{
		Arr: arr,
	}
	qs.QuickSort(0, len(arr)-1)
}

func StackList() {
	fmt.Println("Stack Array....")
	sn := stack.NewStackList()
	sn.Pop()
	sn.Push(1)
	sn.Push(2)
	sn.Push(3)
	sn.Pop()
	sn.Pop()
	sn.Pop()
	sn.Pop()
}

func StackArr() {
	fmt.Println("Stack Array....")

	sn := stack.NewStack()
	sn.Pop()
	sn.Push(1)
	sn.Push(2)
	sn.Push(3)
	sn.Pop()
	sn.Pop()
	sn.Pop()
	sn.Pop()
}

func Queue() {
	sq := make(queue.SliceQ, 0, 5)
	fmt.Printf("sq len:%v, cap:%v\n", len(sq), cap(sq))
	queue.Enqueue(sq, 1)
	queue.Enqueue(sq, 2)
	queue.Traverse(sq)

}

func SinglyList() {
	// Create List
	head := &linkedlist.Node{}
	current := head
	for i := 1; i < 5; i++ {
		nextNode := linkedlist.Create(i)
		current.Next = nextNode
		current = nextNode
	}

	// Traverse List
	linkedlist.Traverse(head)

	// Add nodes
	fmt.Println("Add nodes")
	var addNodes = map[int]int{
		1: 999,
		4: 888,
		2: 22,
		5: 55,
	}
	var err error
	for position, data := range addNodes {
		fmt.Println(data)

		head, err = linkedlist.Add(data, position, head)
		if err != nil {
			fmt.Printf("%v (%v)\n", position, err.Error())
		}

		linkedlist.Traverse(head)
	}

	// Delete nodes
	fmt.Println("Delete nodes")
	var deleteNodes = []int{
		999, 888, 22, 55,
	}
	for _, data := range deleteNodes {
		fmt.Println(data)

		head, err = linkedlist.Delete(data, head)
		if err != nil {
			fmt.Printf("%v (%v)\n", data, err.Error())
		}

		linkedlist.Traverse(head)
	}

	head = linkedlist.Reverse(head)
	linkedlist.Traverse(head)
}

func DoublyList() {
	dl := &linkedlist.DoubleyList{}

	elements := []int{1, 2, 3}
	for k, v := range elements {
		dl.Add(v, k)
		dl.Traversal()
	}

	dl.Add(5, 3)
	dl.Traversal()

	dl.Add(4, 3)
	dl.Traversal()

	dl.Add(0, 0)
	dl.Traversal()

	fmt.Println("Search by Index........")
	atIndex := []int{0, 5, 6}
	for _, v := range atIndex {
		data, err := dl.AtIndex(v)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Printf("[%v]:%v\n", v, data)
		}
	}

	fmt.Println("\nSearch by Value........")
	values := []int{0, 5, 6}
	for _, v := range values {
		index, err := dl.IndexOf(v)
		if err != nil {
			fmt.Printf(err.Error())
		} else {
			fmt.Printf("[%v]:%v\n", index, v)
		}
	}
}

/*
          5         - level = 0
	   /     \
	  4       6     - level = 1
	/   \   /   \
  2         7     9 - level = 2
 /  \            /
1    3          8   - level = 3

Pre-Order:   5->4->2->1->3->6->7->9->8
In-Order:    1->2->3->4->5->7->6->8->9
Post-Order:  1->3->2->4->7->->8->9->6->5
Level-Order: 5->4->6->2->7->9->1->3->8
			 5->4->6->2->9->1->3->7->8->
*/
func BinarySearchTree() {

	fmt.Printf("\n### Binary Search Tree ###")
	bst := &tree.BinarySearchTree{}
	insertData := []int{5, 6, 4, 2, 9, 7, 1, 3, 8}
	for _, data := range insertData {
		fmt.Println("\nAdd: ", data)
		bst.Root = bst.Insert(bst.Root, data)
		//bst.Traverse(bst.Root)
	}

	fmt.Println("\n\n### Tree Traversal ###")
	fmt.Println("PreOrder")
	bst.PreOrderTraverse(bst.Root)
	fmt.Println("\nInOrder")
	bst.InOrderTraverse(bst.Root)
	fmt.Println("\nPostOrder")
	bst.PostOrderTraverse(bst.Root)

	// BFS: LevelOrder Traversal
	fmt.Println("\nLevelOrder")
	bst.LevelOrderTraverse(bst.Root)

	fmt.Println("\n\n### Search Data ###")
	fmt.Printf("9: %v\n", bst.Search(bst.Root, 9))
	fmt.Printf("0: %v\n", bst.Search(bst.Root, 0))
}
