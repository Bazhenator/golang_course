package main

import (
	"container/heap"
	"fmt"
)

type myHeap []int

// Declaring heap's interface
func (h myHeap) Len() int           { return len(h) }
func (h myHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h myHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *myHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *myHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func arraySortDescending(arr []int) {
	h := myHeap(arr)
	//O(n)
	heap.Init(&h)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = heap.Pop(&h).(int)
	}
}

// Example of tested file
func findKthLargest(nums []int, k int) int {
	//Checking array sorting with heap
	arraySortDescending(nums)
	//Cheking if k is out of array's range
	if len(nums) == 0 {
		fmt.Println("null array error was occured")
		return 0
	}
	if k <= len(nums) && k >= 0 {
		return nums[k-1]
	} else {
		fmt.Print("out of range")
	}
	return 0
}

//Heap works, but here another way (optimised quickSort)
/* func quickSort(nums []int, first, last, k int) int {
	if first == last { return nums[first] }

	pivotIndex := partition(nums, first, last)
	rank := pivotIndex - first + 1

	if k == rank {
		return nums[pivotIndex]
	} else if k < rank {
		return quickSort(nums, first, pivotIndex-1, k)
	} else {
		return quickSort(nums, pivotIndex+1, last, k-rank)
	}
}

func partition(nums []int, first, last int) int {
	pivot := nums[last]
	i := first

	for j := first; j < last; j++ {
		if nums[j] >= pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

	nums[i], nums[last] = nums[last], nums[i]
	return i
}

func findKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k)
}
*/
