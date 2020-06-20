package sort

import (
	"fmt"
)

/*
10 | 16 | 8 | 12 | 15 | 6 | 3 | 9 | 5 | 999
i=low                                  j=high
pivot=10 : means find sorted position of 10

PartitionPosition(l,h)
  pivot := a[l]
  i := a[l]
  j := a[h]
  loop until i < j
	loop i++ until a[i] <= pivot
		break
	loop j-- until a[j] > pivot
		break
	if i < j
	swap a[i] with a[j]
  swap a[pivot] with a[j]
  return j // partitioning index

  QuickSort(l,h)
    if (l < h) {
		pIndex = PartitionPosition(l,h)
		QuickSort(l,pIndex)
		QuickSort(pIndex+1,h)
	}

Complexity analysis:
It used recursion & Divide n Conquer Strategy
Best Case:
	In base case, partitions will be in middle i.e median
	So each partiton will be of n/2 of its previous half
	this means n/2/2/2/2/2 ....... until only 1 element is left in each partition
	n/2^k = 1
	n = 2^k
	log(n) = k
	This means in log(n) iterations elements will be sorted, which is also the height of the tree, if we analyse it as
	binary tree having left and right tree with each partition
	Also at each level of tree, the all N elements will be compared/traversed
	So total time will be n*log(n)
Worst Case:
	Only difference will be the height of tree.
	This means partition can be always on one side, with all the elements in one partition, except the pivot element
	So height will be N instead of log(n)
	So total time will be O(n^2)
	Moreover, we can see it as follows:
	comparisons at each level of tree will be n, n-1, n-2 ....... 1 = n(n+1)/2 = O(n^2)

So to achieve Best Case:
  We can always choose middle element of partition as the pivot element

*Stack used will be equal to the height of tree i.e n to log(n)

*/
type QuickSortStruct struct {
	Arr []int
}

func (qs *QuickSortStruct) QuickSort(low, high int) {
	fmt.Printf("QuickSort Index: %v-%v \n", low, high)
	fmt.Println(qs.Arr)
	if low < high {
		pIndex := qs.PartitionLogic(low, high)
		qs.QuickSort(low, pIndex)
		qs.QuickSort(pIndex+1, high)
	}
}

func (qs *QuickSortStruct) PartitionLogic(low, high int) (pIndex int) {
	fmt.Println("Start PartitionLogic:")
	pivot := low
	left := low   //qs.Arr[low]
	right := high //qs.Arr[high]

	fmt.Println("pivot:", qs.Arr[pivot])

	for {
		if left >= right {
			break
		}
		// move from left towards right, until we find element > pivot
		for {
			left++
			if left > right {
				break
			}
			if qs.Arr[left] > qs.Arr[pivot] {
				fmt.Println("found left:", qs.Arr[left])
				break
			}
		}

		// move from right to left until we find element < pivot
		for {
			right--
			if right < left {
				break
			}
			if qs.Arr[right] < qs.Arr[pivot] {
				fmt.Println("found right:", qs.Arr[right])
				break
			}
		}

		// swap the 2 elements found earlier
		if left < right {
			qs.swap(left, right)
		}
		fmt.Println(qs.Arr)
	}

	fmt.Println("found pivot's sorted position")
	qs.swap(right, pivot)

	pIndex = right
	fmt.Println("partitioning index: ", pIndex)
	return
}

func (qs *QuickSortStruct) swap(left, right int) {
	fmt.Printf("swap: left:%v, right:%v\n", qs.Arr[left], qs.Arr[right])
	temp := qs.Arr[left]
	qs.Arr[left] = qs.Arr[right]
	qs.Arr[right] = temp
	return
}

// func PartitionLogic1(arr []int) {
// 	if len(arr) == 1 {
// 		return
// 	}

// 	//j := len(*arr) - 1
// 	pivot := 0
// 	fmt.Printf("arr:%+v\n", arr)

// 	for i := 0; i < len(arr); i++ {

// 		if (arr)[i] > (arr)[pivot] {
// 			for j := len(arr) - 1; j >= i; j-- {
// 				fmt.Printf("%v > %v\n", (arr)[i], (arr)[j])

// 				if (arr)[j] <= (arr)[pivot] {
// 					(arr)[i], (arr)[j] = swap((arr)[i], (arr)[j])
// 					break
// 				}

// 			}
// 		}

// 	}
// 	QuickSort(arr[0 : pivot-1])
// 	QuickSort(arr[pivot+1 : len(arr)-1])
// }
