package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

// Heap's Interface
func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func heapSort(arr []int) {
	h := &IntHeap{}
	heap.Init(h)

	for _, v := range arr {
		heap.Push(h, v)
	}

	for i := 0; i < len(arr); i++ {
		arr[i] = heap.Pop(h).(int)
	}
}

// Example of tested file
func findKthLargest(nums []int, k int) int {
	//Checking Heap sort
	heapSort(nums)
	//Cheking if k is out of array's range
	if k <= len(nums) && k >= 0 {
		return nums[len(nums)-k]
	} else {
		fmt.Print("out of range")
	}
	return 0
}

/*func main() {
	nums := []int{1, 2, 3, 47, 9, 0}
	fmt.Print(findKthLargest(nums, 5))
}*/
