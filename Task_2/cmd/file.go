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

func heapSortDescending(arr []int) {
	h := myHeap(arr)
	heap.Init(&h)
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i] = heap.Pop(&h).(int)
	}
}

// Example of tested file
func findKthLargest(nums []int, k int) int {
	//Checking Heap sort
	heapSortDescending(nums)
	//Cheking if k is out of array's range
	if k <= len(nums) && k >= 0 {
		return nums[k-1]
	} else {
		fmt.Print("out of range")
	}
	return 0
}
