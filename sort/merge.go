package sort

import (
	"fmt"
	"math"
	"sync"
)

/*
2-Way MergeSort: Iterative Method
O(nlog(n)) :
- log(n) pass/traversals
- n elements in each iteration/pass

Pros:
 - Best for Large Lists
 - Linked Lists: without using 3rd list we can only change the pointers of the two lists
 - External Sorting: Instead of copying all data to RAM, Wwe can copy 1 element at a time and copy sorted elements into disk and continue
 - Stable: order of sorting same value elements is maintained

 Cons:
 - Need Extra Space (no inplace sorting): ONLY applicablt to Arrays. In linked lists, no extra space is required
 - Quick Sort can be done InPlace
  - Not for Small size Problem: for  n <= 15, merge sort is slower
 - Recursive: needs log(n) space for stack + n for storing merged list. So total space is O(n+log(n)) = O(n)
*/

type MergeSortStruct struct {
	List []int
}

type MergeStruct struct {
	FirstSortedList  []int
	SecondSortedList []int
}

/*
a1: 1,3,5
a2: 2,4,6
a: 1,2,3,4,5,6

Merging means concatinating 2 sorted lists into a single sorted list
Time: O(m+n), m,n are the number of elements in 2 lists
*/
func (m *MergeStruct) Merge() (mergedList []int) {

	// compare elements of 2 sorted lists
	// append to new list in sorted order
	// until elements of either of the 2 lists get exhausted
	// append the remaining elements of that list to the new list
	var (
		firstIndex, secondIndex         int
		firstExhausted, secondExhausted bool
	)

	for {

		if m.FirstSortedList[firstIndex] < m.SecondSortedList[secondIndex] {
			//	fmt.Println("merge first: ", m.FirstSortedList[firstIndex])
			mergedList = append(mergedList, m.FirstSortedList[firstIndex])
			firstIndex++
		} else {
			//	fmt.Println("merge second: ", m.SecondSortedList[secondIndex])
			mergedList = append(mergedList, m.SecondSortedList[secondIndex])
			secondIndex++
		}

		if firstIndex > len(m.FirstSortedList)-1 {
			firstExhausted = true
			break
		}
		if secondIndex > len(m.SecondSortedList)-1 {
			secondExhausted = true
			break
		}
	}

	if firstExhausted == true {
		if secondIndex <= len(m.SecondSortedList)-1 {
			mergedList = append(mergedList, m.SecondSortedList[secondIndex:len(m.SecondSortedList)]...)
		}
	}

	if secondExhausted == true {
		mergedList = append(mergedList, m.FirstSortedList[firstIndex:len(m.FirstSortedList)]...)
	}
	return
}

/*
list: 1,3,5,2,4
It uses Divide & Conquer Strategy
- Divide RECURSIVELY, the unsorted list of N elements into multiple lists of size 1
- Merge these N lists to form a sorted single list
- Split is done like a binary tree
- Combine is done like POST-ORDER Tree Traversal (Merging Part)
*/

func (ms *MergeSortStruct) MergeSortRoutine(low, high int) (sortedList []int) {
	//fmt.Printf("low:%v, high:%v\n", low, high)
	c1 := make(chan []int)
	c2 := make(chan []int)
	wg := &sync.WaitGroup{}

	if low < high {
		mid := (low + high) / 2

		wg.Add(1)
		go func() {
			defer wg.Done()
			ml1 := ms.MergeSortRoutine(low, mid)
			c1 <- ml1
		}()
		wg.Add(1)
		go func() {
			defer wg.Done()
			ml2 := ms.MergeSortRoutine(mid+1, high)
			c2 <- ml2
		}()

		m := &MergeStruct{
			FirstSortedList:  <-c1,
			SecondSortedList: <-c2,
		}
		wg.Wait()

		sortedList = m.Merge()
	} else {
		sortedList = ms.List[low : high+1]
	}
	return
}

func (ms *MergeSortStruct) MergeSort(low, high int) (sortedList []int) {
	//fmt.Printf("low:%v, high:%v\n", low, high)

	if low < high {
		mid := (low + high) / 2
		ml1 := ms.MergeSort(low, mid)
		ml2 := ms.MergeSort(mid+1, high)

		m := &MergeStruct{
			FirstSortedList:  ml1,
			SecondSortedList: ml2,
		}
		sortedList = m.Merge()
	} else {
		sortedList = ms.List[low : high+1]
	}
	return
}

/*
iteration-1
[0:1], [1:2]
[2:3], [3:4]
[4:5], [5:6]
[6:7], [7:8]
iteration-2
[0:2], [2:4]
[4:6], [6:8]
iteration-3
[0:4], [4:8]

*/
func (ms *MergeSortStruct) TwoWayMerge() (sortedList []int) {

	fmt.Println(ms.List)
	var start int
	for j := 1; j < 4; j++ {
		factor := int(math.Pow(2, float64(j)))
		sortedList = nil
		for i := 0; i < len(ms.List)/factor; i++ {

			//fmt.Printf("i:%v, j:%v\n", i, j)
			if i == 0 {
				start = i
			}
			l1 := start
			h1 := start + factor/2
			l2 := h1
			h2 := l2 + factor/2

			//fmt.Printf("L1[%v:%v], L2[%v:%v]\n", l1, h1, l2, h2)
			m1 := &MergeStruct{
				FirstSortedList:  ms.List[l1:h1],
				SecondSortedList: ms.List[l2:h2],
			}
			mergerList := m1.Merge()
			sortedList = append(sortedList, mergerList...)

			start = h2
		}
		//ms.List = nil
		ms.List = sortedList
		fmt.Println(sortedList, ms.List)
	}

	return
}

type MergeListNode struct {
	Data int
	Next *MergeListNode
}

type MergeList struct {
	Head *MergeListNode
}

func NewMergeList() *MergeList {
	return &MergeList{}
}
func (ml *MergeList) CreateNode(data int) *MergeListNode {
	//fmt.Println("CreateNode")
	return &MergeListNode{Data: data}
}

func (ml *MergeList) InsertNode(data int) {
	//fmt.Println("InsertNode")

	newNode := ml.CreateNode(data)
	current := ml.Head
	for {
		// if first element
		if ml.Head == nil {
			//	fmt.Println("First Node")
			ml.Head = newNode
			return
		}
		// append node to the end of the list
		if current.Next == nil {
			current.Next = newNode
			return
		}
		current = current.Next
	}

}

func (ml *MergeList) TraverseList() (err error) {
	if ml.Head == nil {
		err = fmt.Errorf("Empty list")
		return
	}
	current := ml.Head
	for {
		fmt.Printf("%v->", current.Data)
		if current.Next == nil {
			return
		}
		current = current.Next
	}
}

func MergeLinkedLists(firstHead, secondHead *MergeListNode) *MergeListNode {

	var newHead, newCurrent *MergeListNode
	currFirst := firstHead
	currSecond := secondHead

	for {

		if currFirst.Data < currSecond.Data {
			if newHead == nil {
				newHead = currFirst
				newCurrent = newHead
			} else {
				newCurrent.Next = currFirst
				newCurrent = currFirst
			}
			currFirst = currFirst.Next

		} else {
			if newHead == nil {
				newHead = currSecond
				newCurrent = newHead
			} else {
				newCurrent.Next = currSecond
				newCurrent = currSecond
			}
			currSecond = currSecond.Next
		}

		if currFirst == nil || currSecond == nil {
			break
		}

	}

	if currFirst == nil {
		newCurrent.Next = currSecond
	}
	if currSecond == nil {
		newCurrent.Next = currFirst
	}

	return newHead
}
